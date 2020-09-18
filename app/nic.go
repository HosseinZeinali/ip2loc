package app

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/HosseinZeinali/ip2loc/dto"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Nic struct {
}

type NicTitle string

const (
	afrinic       = "afrinic"
	lacnic        = "lacnic"
	arin          = "arin"
	apnic         = "apnic"
	ripe          = "ripe"
	filesBasePath = "./data"
)

var nicsList = []string{
	lacnic,
	afrinic,
	arin,
	apnic,
	ripe,
}

var nicUrlMapper = map[string]string{
	afrinic: "https://ftp.afrinic.net/pub/stats/afrinic/delegated-afrinic-latest",
	lacnic:  "https://ftp.lacnic.net/pub/stats/lacnic/delegated-lacnic-latest",
	arin:    "https://ftp.arin.net/pub/stats/arin/delegated-arin-extended-latest",
	apnic:   "https://ftp.apnic.net/pub/stats/apnic/delegated-apnic-latest",
	ripe:    "https://ftp.ripe.net/ripe/stats/delegated-ripencc-latest",
}

func (nic *Nic) downloadNicDataFileByNicTitle(nicTitle string) error {
	_, ok := nicUrlMapper[nicTitle]
	var doesExist = false
	if ok {
		doesExist = true
	}

	if !doesExist {
		return errors.New(fmt.Sprintf("unefined index %s", nicTitle))
	}

	err := downloadFile(filesBasePath+"/"+nicTitle, nicUrlMapper[nicTitle])

	return err
}

func (nic *Nic) downloadNicDataFileMd5ByNicTitle(nicTitle string) error {
	_, ok := nicUrlMapper[nicTitle]
	var doesExist = false
	if ok {
		doesExist = true
	}

	if !doesExist {
		return errors.New(fmt.Sprintf("unefined index %s", nicTitle))
	}

	err := downloadFile(fmt.Sprintf("%s.md5", filesBasePath+"/"+nicTitle), fmt.Sprintf("%s.md5", nicUrlMapper[nicTitle]))

	return err
}

func (nic *Nic) DownloadNicData() error {
	for _, nicTitle := range nicsList {
		if err := nic.downloadNicDataFileByNicTitle(nicTitle); err != nil {
			panic(err)
		}
		if err := nic.downloadNicDataFileMd5ByNicTitle(nicTitle); err != nil {
			panic(err)
		}
	}

	return nil
}

func (nic *Nic) CheckForChangeByNicTitle(nicTitle string) (bool, error) {
	_, ok := nicUrlMapper[nicTitle]
	var doesExist = false
	if ok {
		doesExist = true
	}

	if !doesExist {
		return false, errors.New(fmt.Sprintf("unefined index %s", nicTitle))
	}

	file, err := os.Open(fmt.Sprintf("%s.md5", filesBasePath+"/"+nicTitle))
	if err != nil {
		//log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	systemMd5 := ""
	for scanner.Scan() {
		systemMd5 = strings.Trim(scanner.Text(), "\t \n")
	}

	internetMd5 := strings.Trim(httpGet(fmt.Sprintf("%s.md5", nicUrlMapper[nicTitle])), "\t \n")

	return systemMd5 != internetMd5, nil
}

func (nic *Nic) CheckForChange() (bool, error) {
	for _, nicTitle := range nicsList {
		isChanged, err := nic.CheckForChangeByNicTitle(nicTitle)

		if err != nil {
			return false, err
		}
		if isChanged {
			return isChanged, nil
		}
	}

	return false, nil
}

func (nic *Nic) GetNicRecordsByNicTitle(nicTitle string) (<-chan *dto.IpDto, error) {
	_, ok := nicUrlMapper[nicTitle]
	var doesExist = false
	if ok {
		doesExist = true
	}

	if !doesExist {
		return nil, errors.New(fmt.Sprintf("unefined index %s", nicTitle))
	}

	reader, err := readlines(fmt.Sprintf("%s/%s", filesBasePath, nicTitle))
	if err != nil {
		log.Fatal(err)
	}

	chnl := make(chan *dto.IpDto)
	go func() {
		for line := range reader {
			if getNicFileLineType(line) == "ipv4" {
				ip := Nic2Ipv4Dto(line)
				chnl <- ip
			}
			if getNicFileLineType(line) == "ipv6" {
				ip := Nic2Ipv6Dto(line)
				chnl <- ip
			}
		}
		close(chnl)
	}()

	return chnl, nil
}

func (nic *Nic) GetNicRecords() (<-chan *dto.IpDto, error) {
	chnl := make(chan *dto.IpDto)
	go func() {
		for _, nicTitle := range nicsList {
			reader, err := nic.GetNicRecordsByNicTitle(nicTitle)
			if err != nil {
				log.Fatal(err)
			}
			for ip := range reader {
				chnl <- ip
			}
		}
		close(chnl)
	}()

	return chnl, nil
}

func downloadFile(filepath string, url string) error {
	client := http.Client{
		Timeout: -1,
	}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func readlines(path string) (<-chan string, error) {
	fobj, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fobj)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	chnl := make(chan string)
	go func() {
		for scanner.Scan() {
			chnl <- scanner.Text()
		}
		close(chnl)
	}()

	return chnl, nil
}

func getNicFileLineType(line string) string {
	ipSlice := strings.Split(line, "|")
	if (len(ipSlice) == 7 || len(ipSlice) == 8) && (ipSlice[2] == "ipv4" || ipSlice[2] == "ipv6") {
		if ipSlice[2] == "ipv4" {
			return "ipv4"
		} else {
			return "ipv6"
		}
	}
	return ""
}

func httpGet(url string) string {
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(content)
}

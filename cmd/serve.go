package cmd

import (
	"fmt"
	apiHttp "github.com/HosseinZeinali/ip2loc/api/http"
	"github.com/HosseinZeinali/ip2loc/app"
	"github.com/HosseinZeinali/ip2loc/model"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"time"
)

func Serve() {
	ticker := time.NewTicker(10 * time.Minute)
	app, _ := app.New()
	ctx := app.NewContext()

	if !ctx.Database.DoesTableExist("tables") {
		ctx.Database.CreateTableTable()
	}

	go func() {
		for ; true; <-ticker.C {
			updateIpsIfNeeded(ctx)
		}
	}()
	api1, _ := apiHttp.New(app)
	router := mux.NewRouter()
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)
	s := &http.Server{
		Addr:        ":8081",
		Handler:     cors(router),
		ReadTimeout: 2 * time.Minute,
	}
	api1.Init(router.PathPrefix("/api").Subrouter())
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Error(err)
	}
}

func updateIpsIfNeeded(ctx *app.Context) {
	isChanged, _ := ctx.Nic.CheckForChange()
	if isChanged {
		fmt.Println("The database should be updated")
	} else {
		fmt.Println("The database is already up-to-date.")
	}
	if isChanged {
		ipv4Table := "ipv4s_" + RandStringRunes(8)
		ipv6Table := "ipv6s_" + RandStringRunes(8)
		fmt.Println("new ipv4 table name:" + ipv4Table)
		fmt.Println("new ipv6 table name:" + ipv6Table)
		fmt.Println("donwloding nic data")
		ctx.Nic.DownloadNicData()
		ctx.Database.CreateIpv4Table(ipv4Table)
		ctx.Database.CreateIpv6Table(ipv6Table)
		reader, _ := ctx.Nic.GetNicRecords()
		i := 1
		j := 1
		var ipv4s []*model.Ipv4
		var ipv6s []*model.Ipv6
		for ip := range reader {
			if ip.IpType == "ipv4" {
				ipv4 := app.IpDto2Ipv4(ip)
				ipv4s = append(ipv4s, ipv4)
				if i%100 == 0 {
					err := ctx.Database.CreateBatchIpv4s(ipv4Table, ipv4s)
					if err != nil {
						fmt.Println("error occurred inserting 100 ipv4 records")
					} else {
						fmt.Println("100 ipv4 records inserted successfully")
					}
					ipv4s = nil
				}
				i++
			}
			if ip.IpType == "ipv6" {
				ipv6 := app.IpDto2Ipv6(ip)
				ipv6s = append(ipv6s, ipv6)
				if j%100 == 0 {
					err := ctx.Database.CreateBatchIpv6s(ipv6Table, ipv6s)
					if err != nil {
						fmt.Println("error occurred inserting 100 ipv6 records")
					} else {
						fmt.Println("100 ipv6 records inserted successfully")
					}
					ipv6s = nil
				}
				j++
			}
		}
		ctx.Database.CreateBatchIpv4s(ipv4Table, ipv4s)
		ctx.Database.CreateBatchIpv6s(ipv6Table, ipv6s)
		ipv4T := model.NewTable(ipv6Table, "", time.Now(), "ipv6")
		ipv6T := model.NewTable(ipv4Table, "", time.Now(), "ipv4")
		ctx.Database.CreateTable(ipv4T)
		ctx.Database.CreateTable(ipv6T)
		ctx.Database.DefaultIpv6Table = ipv6Table
		ctx.Database.DefaultIpv4Table = ipv4Table
		fmt.Println("Database updated successfully.")
	}
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

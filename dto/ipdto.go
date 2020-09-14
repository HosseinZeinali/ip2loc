package dto

import "time"

type IpDto struct {
	ID      uint64
	IpType  string
	IpStart string
	IpEnd   string
	IpCount int
	Date    time.Time
	Status  string
	Cc      string
}

func NewIpDto(ipType string, ipStart string, ipEnd string, ipCount int, date time.Time, status string, cc string) *IpDto {
	ipDto := new(IpDto)
	ipDto.ID = 0
	ipDto.IpType = ipType
	ipDto.IpStart = ipStart
	ipDto.IpEnd = ipEnd
	ipDto.IpCount = ipCount
	ipDto.Date = date
	ipDto.Status = status
	ipDto.Cc = cc

	return ipDto
}

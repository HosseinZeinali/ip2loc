package model

import (
	"math/big"
	"time"
)

type Ipv6 struct {
	ID      int       `db:"id"`
	IpType  string    `db:"ip_type"`
	IpStart string    `db:"ip_start"`
	IpEnd   string    `db:"ip_end"`
	IpCount int       `db:"ip_count"`
	Date    time.Time `db:"date"`
	Status  string    `db:"status"`
	Cc      string    `db:"cc"`
}

/*type Ipv62 struct {
	ID int `db:"id"`
	IpType string `db:"ip_type"`
	IpStart string `db:"ip_start"`
	IpEnd string `db:"ip_end"`
	IpCount int `db:"ip_count"`
	Date time.Time `db:"date"`
	Status string `db:"status"`
	Cc string `db:"cc"`
}*/

func NewIpv6(ipStart *big.Int, ipEnd *big.Int, ipCount int, date time.Time, status string, cc string) *Ipv6 {
	ipv6 := new(Ipv6)
	ipv6.IpStart = ipStart.String()
	ipv6.IpEnd = ipEnd.String()
	ipv6.IpCount = ipCount
	ipv6.Date = date
	ipv6.Status = status
	ipv6.Cc = cc

	return ipv6
}

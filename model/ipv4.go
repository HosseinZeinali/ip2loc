package model

import "time"

type Ipv4 struct {
	ID      int       `db:"id"`
	IpType  string    `db:"ip_type"`
	IpStart uint32    `db:"ip_start"`
	IpEnd   uint32    `db:"ip_end"`
	IpCount int       `db:"ip_count"`
	Date    time.Time `db:"date"`
	Status  string    `db:"status"`
	Cc      string    `db:"cc"`
}

func NewIpv4(ipStart uint32, ipEnd uint32, ipCount int, date time.Time, status string, cc string) *Ipv4 {
	ipv4 := new(Ipv4)
	ipv4.IpStart = ipStart
	ipv4.IpEnd = ipEnd
	ipv4.IpCount = ipCount
	ipv4.Date = date
	ipv4.Status = status
	ipv4.Cc = cc

	return ipv4
}

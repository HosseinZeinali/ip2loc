package sqlx

import (
	"fmt"
	"github.com/HosseinZeinali/ip2loc/model"
	"strconv"
)

func (db *Database) CreateIpv4(tableName string, ip *model.Ipv4) error {
	ipState := fmt.Sprintf(`INSERT INTO %s (ip_type, ip_start, ip_end, ip_count, date, status, cc)
				VALUES (:type, :ip_start, :ip_end, :ip_count, :date, :status, :cc)`, tableName)
	m := map[string]interface{}{
		"type":     ip.IpType,
		"ip_start": ip.IpStart,
		"ip_end":   ip.IpEnd,
		"ip_count": strconv.Itoa(ip.IpCount),
		"date":     ip.Date.Format("2 Jan 2006 15:04:05"),
		"status":   ip.Status,
		"cc":       ip.Cc,
	}
	_, err := db.NamedExec(ipState, m)

	return err
}

func (db *Database) CreateBatchIpv4s(tableName string, ips []*model.Ipv4) error {
	tx := db.MustBegin()
	for _, ip := range ips {
		ipState := fmt.Sprintf(`INSERT INTO %s (ip_type, ip_start, ip_end, ip_count, date, status, cc)
				VALUES (:type, :ip_start, :ip_end, :ip_count, :date, :status, :cc)`, tableName)
		m := map[string]interface{}{
			"type":     ip.IpType,
			"ip_start": ip.IpStart,
			"ip_end":   ip.IpEnd,
			"ip_count": strconv.Itoa(ip.IpCount),
			"date":     ip.Date.Format("2 Jan 2006 15:04:05"),
			"status":   ip.Status,
			"cc":       ip.Cc,
		}
		tx.NamedExec(ipState, m)
	}
	err := tx.Commit()

	return err
}

func (db *Database) GetIpv4RecordByIpv4(tableName string, intIp uint32) (*model.Ipv4, error) {
	m := map[string]interface{}{
		"ip": intIp,
	}

	rows, err := db.NamedQuery(fmt.Sprintf(`SELECT * FROM %s WHERE ip_start <= :ip AND ip_end > :ip LIMIT 1;`, tableName), m)
	var ip model.Ipv4
	for rows.Next() {
		err = rows.StructScan(&ip)
	}

	return &ip, err
}

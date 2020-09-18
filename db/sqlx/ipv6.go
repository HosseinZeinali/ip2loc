package sqlx

import (
	"fmt"
	"github.com/HosseinZeinali/ip2loc/model"
	"math/big"
	"strconv"
)

func (db *Database) CreateIpv6(tableName string, ip *model.Ipv6) error {
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

func (db *Database) CreateBatchIpv6s(tableName string, ips []*model.Ipv6) error {
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

func (db *Database) GetIpv6RecordByIpv6(tableName string, intIp *big.Int) (*model.Ipv6, error) {
	m := map[string]interface{}{
		"ip": intIp.String(),
	}

	ipState := fmt.Sprintf(`SELECT * FROM %s WHERE ip_start <= :ip AND ip_end > :ip LIMIT 1;`, tableName)
	rows, err := db.NamedQuery(ipState, m)
	var ip model.Ipv6
	for rows.Next() {
		err = rows.StructScan(&ip)
	}

	return &ip, err
}

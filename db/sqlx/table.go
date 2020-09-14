package sqlx

import (
	"fmt"
	"github.com/HosseinZeinali/ip2loc/model"
)

func (db *Database) CreateTable(table *model.Table) error {
	tableState := `INSERT INTO tables (name, extra, date, type)
					VALUES (:name, :extra, :date, :type)`
	m := map[string]interface{}{
		"name":  table.Name,
		"extra": table.Extra,
		"date":  table.Date.Format("2 Jan 2006 15:04:05"),
		"type":  table.Type,
	}
	_, err := db.NamedExec(tableState, m)

	return err
}

func (db *Database) GetAllTables() ([]*model.Table, error) {
	tableState := fmt.Sprintf(`SELECT * FROM tables ORDER BY id DESC; `)
	rows, err := db.Queryx(tableState)
	var tables []*model.Table
	for rows.Next() {
		var table model.Table
		err = rows.StructScan(&table)
		if err != nil {
			return nil, err
		}
		tables = append(tables, &table)
	}

	return tables, err

}

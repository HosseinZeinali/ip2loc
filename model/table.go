package model

import "time"

type Table struct {
	ID    int       `db:"id"`
	Name  string    `db:"name"`
	Date  time.Time `db:"date"`
	Extra string    `db:"extra"`
	Type  string    `db:"type"`
}

func NewTable(name string, extra string, date time.Time, tableType string) *Table {
	table := new(Table)
	table.Name = name
	table.Extra = extra
	table.Date = date
	table.Type = tableType

	return table
}

package app

import db "github.com/HosseinZeinali/ip2loc/db/sqlx"

type Context struct {
	Database *db.Database
	Nic      *Nic
}

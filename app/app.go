package app

import (
	db "github.com/HosseinZeinali/ip2loc/db/sqlx"
)

type App struct {
	Config   *db.Config
	Database *db.Database
}

func (a *App) NewContext() *Context {
	nic := new(Nic)
	return &Context{
		Database: a.Database,
		Nic:      nic,
		Logger:   NewLogger(),
	}
}

func New() (app *App, err error) {
	app = &App{}
	dbConfig, err := db.InitConfig()
	if err != nil {
		return nil, err
	}
	app.Database, err = db.New(dbConfig)
	tables, _ := app.Database.GetAllTables()
	for _, table := range tables {
		if app.Database.DefaultIpv4Table != "" && app.Database.DefaultIpv6Table != "" {
			break
		}
		if table.Type == "ipv6" && app.Database.DefaultIpv6Table == "" {
			app.Database.DefaultIpv6Table = table.Name
		}
		if table.Type == "ipv4" && app.Database.DefaultIpv4Table == "" {
			app.Database.DefaultIpv4Table = table.Name
		}
	}

	if err != nil {
		return nil, err
	}

	return app, err
}

func (a *App) Close() error {
	return a.Database.Close()
}

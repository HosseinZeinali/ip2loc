package sqlx

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	*sqlx.DB
	Config           *Config
	DefaultIpv4Table string
	DefaultIpv6Table string
}

func New(config *Config) (*Database, error) {
	db, err := sqlx.Connect("postgres", config.DatabaseURI)
	if err != nil {
		log.Fatalln(err)
	}

	return &Database{db, config, "", ""}, nil
}

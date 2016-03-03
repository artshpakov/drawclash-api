package app

import (
	"github.com/artshpakov/drawclash/config"
	"gopkg.in/pg.v3"
)

var db *pg.DB

func GetDB() *pg.DB {
	if db == nil {
		return pg.Connect(&pg.Options{Addr: config.DB_ADDR, User: config.DB_USER, Database: config.DB_NAME})
	}
	return db
}

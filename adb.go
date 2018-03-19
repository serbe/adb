package adb

import (
	"log"

	"github.com/go-pg/pg"
)

// ADB - structure for interacting with the database
type ADB struct {
	db *pg.DB
}

// InitDB - initializing the connection to the database
func InitDB(database, addr, user, password string) *ADB {
	a := new(ADB)
	db := pg.Connect(&pg.Options{
		Database: database,
		Addr:     addr,
		User:     user,
		Password: password,
	})
	_, err := db.Exec("SELECT NULL LIMIT 0")
	if err != nil {
		log.Fatal("InitDB error: ", err)
	}
	a.db = db
	return a
}

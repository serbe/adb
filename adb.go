package adb

import (
	"log"

	"github.com/go-pg/pg"
)

var debug bool

type ADB struct {
	db *pg.DB
}

func InitDB(database, addr, user, password string) (*ADB, error) {
	a := new(ADB)
	db := pg.Connect(&pg.Options{
		Database: database,
		Addr:     addr,
		User:     user,
		Password: password,
	})
	_, err := db.Exec("SELECT NULL LIMIT 0")
	a.db = db
	return a, err
}

func (a *ADB) UseDebug(value bool) {
	debug = value
}

func chkErr(msg string, err error) {
	if debug {
		log.Println("error:", msg, err)
	}
}

package adb

import (
	"log"

	"github.com/go-pg/pg"
)

var debug bool

type ADB struct {
	db *pg.DB
}

func InitDB(addr, user, password string) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
	})
	_, err := db.Exec("SELECT NULL LIMIT 0")
	return db, err
}

func (a *ADB) UseDebug(value bool) {
	debug = value
}

func chkErr(msg string, err error) {
	if debug {
		log.Println("error:", msg, err)
	}
}

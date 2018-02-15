package adb

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

var useShowError bool

type ADB struct {
	db *pg.DB
}

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
		log.Println("InitDB error: ", err)
		os.Exit(1)
	}
	a.db = db
	return a
}

func chkErr(msg string, err error) {
	if useShowError && err != nil {
		log.Println("error:", msg, err)
	}
}

package adb

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	logErrors bool
)

// DB client pool
type DB struct {
	Pool *pgxpool.Pool
}

// InitDB - initializing the connection to the database
func InitDB(dbURL string) DB {
	var db DB
	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Printf("Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	_, err = pool.Exec(context.Background(), "SELECT NULL LIMIT 0")
	if err != nil {
		log.Fatal("InitDB error: ", err)
	}
	db.Pool = pool
	return db
}

func errmsg(str string, err error) {
	if logErrors {
		log.Println("Error in", str, err)
	}
}

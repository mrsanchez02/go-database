package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://edteam:edteam@localhost:7530/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping: %v", err)
		}

		fmt.Println("Connected to Postgres")
	})

}

func NewMySQLDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "leandrosc:leandrosc@tcp(localhost:3306)/godb")
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping: %v", err)
		}
		fmt.Println("Connected to Mysql")
	})
}

// Pool return a unique instancie of db
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}

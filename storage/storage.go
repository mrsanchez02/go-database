package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/mrsanchez02/go-database/pkg/product"
)

var (
	db   *sql.DB
	once sync.Once
)

// Driver of storage
type Driver string

const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
func New(d Driver) {
	switch d {
	case "MYSQL":
		newMySQLDB()
	case "POSTGRES":
		newPostgresDB()
	default:

	}
}

func newPostgresDB() {
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

func newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "leandrosc:leandrosc@tcp(localhost:3306)/godb?parseTime=true")
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

// DAOProduct factory for product.Storage
func DAOProduct(driver Driver) (product.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySQL:
		return newMySQLProduct(db), nil
	default:
		return nil, fmt.Errorf("Driver not implemented")
	}
}

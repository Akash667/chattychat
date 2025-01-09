package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type dbObject struct {
	db *sql.DB
}

func NewDatabase() *dbObject {
	connStr := "user=root password=mysecretpassword host=localhost port=5433 sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Connection to database failed")
	}

	return &dbObject{db: db}

}

func (d *dbObject) Close() {
	d.db.Close()
}

func (d *dbObject) GetDB() *sql.DB {
	return d.db
}

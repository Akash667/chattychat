package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type dbObject struct {
	db *sql.DB
}

func NewDatabase() *dbObject {

	// connStr := "user=root password=mysecretpassword host=localhost port=5433 sslmode=disable"
	_, filePath, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filePath)
	dbPath := filepath.Join(dir, "basechat.db")
	fmt.Println("Database path:", dbPath)

	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		fmt.Println(err)
	}

	sqlBytes, err := os.ReadFile(dir + "/migrations/20241212192608_add_users_table.up.sql")
	if err != nil {
		fmt.Println("migration failed", err)
	}

	sqlStmt := string(sqlBytes)
	log.Println("creating database", sqlStmt)
	_, err = db.Exec(sqlStmt)

	if err != nil {
		fmt.Println(err)
	}

	return &dbObject{db: db}

}

func (d *dbObject) Close() {
	d.db.Close()
}

func (d *dbObject) GetDB() *sql.DB {
	return d.db
}

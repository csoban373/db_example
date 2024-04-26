package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username = "root"
	password = ""
	host     = "127.0.0.1"
	port     = "3306"
	dbName   = "people"
)

func SetupDB() (*sql.DB, error) {
	fmt.Println("Database connection...")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		username,
		password,
		host,
		port,
		dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestDB(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		fmt.Println("failed to ping!")
		return err
	}
	fmt.Println("Succesful ping!")
	return nil
}

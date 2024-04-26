package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	username = "root"
	password = ""
	host     = "127.0.0.1"
	port     = "3306"
	dbName   = "people"
)

func SetupDB() (*gorm.DB, error) {
	fmt.Println("Database connection...")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		username,
		password,
		host,
		port,
		dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestDB(dbConn *gorm.DB) error {
	db, err := dbConn.DB()
	if err != nil {
		fmt.Println("failed to get db!")
		return err
	}
	if err = db.Ping(); err != nil {
		fmt.Println("failed to ping!")
		return err
	}
	fmt.Println("Succesful ping!")
	return nil
}

func CloseDB(dbConn *gorm.DB) error {
	db, err := dbConn.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

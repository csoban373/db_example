package main

import (
	"fmt"

	"github.com/csoban373/db_example/db"
)

func main() {
	dbConn, err := db.SetupDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.CloseDB(dbConn)

	err = db.TestDB(dbConn)
	if err != nil {
		fmt.Println(err)
	}
}

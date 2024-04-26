package main

import (
	"fmt"

	"github.com/csoban373/db_example/db"
	"github.com/csoban373/db_example/handler"
)

func main() {
	dbConn, err := db.SetupDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbConn.Close()

	err = db.TestDB(dbConn)
	if err != nil {
		fmt.Println(err)
		return
	}

	allPeople, err := handler.GetAllPeople(dbConn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(allPeople)
}

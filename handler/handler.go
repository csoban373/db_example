package handler

import (
	"database/sql"
	"fmt"

	"github.com/csoban373/db_example/model"
)

func GetAllPeople(dbConn *sql.DB) ([]model.Person, error) {
	allPeople := make([]model.Person, 0)

	results, err := dbConn.Query("SELECT * FROM person")
	if err != nil {
		fmt.Println("failed to find all people!")
		return allPeople, err
	}
	defer results.Close()

	for results.Next() {
		var person model.Person
		err = results.Scan(&person.ID, &person.Name, &person.Age)
		if err != nil {
			return allPeople, err
		}
		allPeople = append(allPeople, person)
	}

	return allPeople, nil
}

package handler

import (
	"fmt"

	"github.com/csoban373/db_example/model"
	"gorm.io/gorm"
)

func GetAllPeople(dbConn *gorm.DB) ([]model.Person, error) {
	var result []model.Person
	if err := dbConn.Find(&result).Error; err != nil {
		fmt.Println("failed to find all people!")
		return result, err
	}
	return result, nil
}

package database

import (
	"go-crud/data"
	"log"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Migrate() {
	Connector.AutoMigrate(&data.Person{})
}

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	log.Println("Connection was sucessful")
	return nil
}

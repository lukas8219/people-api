package database

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

func Open() {
	config := Config {
		User: "lucas",
		Password: "12345678",
		DB: "go",
		Host: "localhost",
		Port: "3306",
	}

	connectionString := GetConnectionString(config)
	fmt.Println(connectionString)
	err := Connect(connectionString)

	if err != nil {
		panic(err.Error())
	}
}

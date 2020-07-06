package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {

	username := "root"
	password := "p@ssw0rd"
	dbName := "testGorm"
	dbHost := "localhost"

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	database, err := gorm.Open("mysql", dbUri)

	if err != nil {
		fmt.Print(err)
	}

	database.AutoMigrate(&Book{})

	DB = database
}

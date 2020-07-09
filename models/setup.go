package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {

	username := "root"
	password := "p@ssw0rd"
	dbName := "test-gorm"
	dbHost := "localhost"
	dbPort := "3306"
	dbType := "mysql"

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, dbHost, dbPort, dbName)
	fmt.Println(dbURI)

	database, err := gorm.Open(dbType, dbURI)

	if err != nil {
		fmt.Print(err)
	}

	DB = database
}

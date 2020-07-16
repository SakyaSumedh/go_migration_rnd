package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/SakyaSumedh/go_migration_rnd/controllers"
	"github.com/SakyaSumedh/go_migration_rnd/models"
)

func main() {
	r := gin.Default()

	m, err := migrate.New(
		"file://migrations/",
		"mysql://root:p@ssw0rd@tcp(localhost:3306)/test-gorm")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Migration up")
	m.Steps(1000)

	models.ConnectDatabase()

	r.GET("/publications", controllers.FindPublications)
	r.POST("/publications", controllers.CreatePublication)
	r.GET("/publications/:id", controllers.FindPublication)
	r.PUT("/publications/:id", controllers.UpdatePublication)
	r.DELETE("/publications/:id", controllers.DeletePublication)

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}

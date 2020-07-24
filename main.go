package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/SakyaSumedh/go_migration_rnd/controllers"
	"github.com/SakyaSumedh/go_migration_rnd/models"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	DBType   string
}

func buildDBConfig(host, port, user, name, password string, dbType string) *DBConfig {
	dbConfig := DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		DBName:   name,
		Password: password,
		DBType:   dbType,
	}
	return &dbConfig
}

func dbURL(dbConfig *DBConfig) string {
	if dbConfig.DBType == "cloudsql" {
		return fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.DBName,
		)
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func main() {
	r := gin.Default()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbType := os.Getenv("DB_TYPE")

	fmt.Println("--- Connecting Migrations ---")
	dbConfig := buildDBConfig(dbHost, dbPort, dbUser, dbName, dbPassword, dbType)
	dbURL := dbURL(dbConfig)

	m, err := migrate.New("file://api/migrations/", dbType+"://"+dbURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("--- Running Migration ---")
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

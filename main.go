package main

import (
	"github.com/gin-gonic/gin"

	"github.com/SakyaSumedh/testGorm/models"
)

func main() {
	r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	// })

	models.ConnectDatabase()

	r.Run()
}

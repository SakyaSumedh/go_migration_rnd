package controllers

import (
	"net/http"

	"github.com/SakyaSumedh/go_migration_rnd/models"
	"github.com/gin-gonic/gin"
)

// CreateBookInput struct for creating book object
type CreateBookInput struct {
	Title         string `json:"title" binding:"required"`
	Author        string `json:"author" binding:"required"`
	PublicationID uint   `json:"publication_id" binding:"required"`
}

// UpdateBookInput struct for creating book object
type UpdateBookInput struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublicationID uint   `json:"publication_id"`
}

// FindBooks GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var publication models.Publication
	if err := models.DB.Where("id = ?", input.PublicationID).First(&publication).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Publication Not Found!"})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author, PublicationID: publication.ID}
	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// FindBook GET /books/:id
func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook PUT /books/:id
func UpdateBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var publication models.Publication
	if err := models.DB.Where("id = ?", input.PublicationID).First(&publication).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Publication Not Found!"})
		return
	}

	models.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook DELETE /books/:id
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)
	c.JSON(http.StatusNoContent, gin.H{})
}

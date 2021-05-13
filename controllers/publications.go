package controllers

import (
	"net/http"

	"github.com/SakyaSumedh/go_migration_rnd/models"
	"github.com/gin-gonic/gin"
)

// CreatePublicationInput struct for creating publication object
type CreatePublicationInput struct {
	Title string `json:"title" binding:"required"`
}

// UpdatePublicationInput struct for creating publication object
type UpdatePublicationInput struct {
	Title string `json:"title"`
}

// FindPublications GET /publications
// Get all publications
func FindPublications(c *gin.Context) {
	var publications []models.Publication
	models.DB.Find(&publications)

	c.JSON(http.StatusOK, gin.H{"data": publications})
}

// CreatePublication POST /publications
// Create new publication
func CreatePublication(c *gin.Context) {
	// Validate input
	var input CreatePublicationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create publication
	publication := models.Publication{Title: input.Title}
	models.DB.Create(&publication)

	c.JSON(http.StatusOK, gin.H{"data": publication})
}

// FindPublication GET /publications/:id
func FindPublication(c *gin.Context) {
	var publication models.Publication

	if err := models.DB.Where("id = ?", c.Param("id")).First(&publication).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": publication})
}

// UpdatePublication PUT /publications/:id
func UpdatePublication(c *gin.Context) {
	var publication models.Publication

	if err := models.DB.Where("id = ?", c.Param("id")).First(&publication).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdatePublicationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&publication).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": publication})
}

// DeletePublication DELETE /publications/:id
func DeletePublication(c *gin.Context) {
	var publication models.Publication
	if err := models.DB.Where("id = ?", c.Param("id")).First(&publication).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&publication)
	c.JSON(http.StatusNoContent, gin.H{})
}

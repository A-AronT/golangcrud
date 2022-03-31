package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

func FindAuthors(c *gin.Context) {
	var authors []models.Author
	models.DB.Find(&authors)

	c.JSON(http.StatusOK, gin.H{"data": authors})
}

func FindAuthor(c *gin.Context) {
	// Get model if exist
	var author models.Author
	if err := models.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": author})
}


type CreateAuthorInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateAuthor(c *gin.Context) {
	// Validate input
	var input CreateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create author
	author := models.Author{Name: input.Name}
	models.DB.Create(&author)

	c.JSON(http.StatusOK, gin.H{"data": author})
}
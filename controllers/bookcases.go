package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

type CreateBookcaseInput struct {
	Title  string `json:"title" binding:"required"`
}

func CreateBookCase(c *gin.Context) {
	// Validate input
	var input CreateBookcaseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create bookcase
	bookcase := models.Bookcase{Title: input.Title}
	models.DB.Create(&bookcase)

	c.JSON(http.StatusOK, gin.H{"data": bookcase})
}
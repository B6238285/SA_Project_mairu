package controller

import (
	"net/http"

	"github.com/B6238285/week5/entity"
	"github.com/gin-gonic/gin"
)

// POST /resolutions
func CreateBookType(c *gin.Context) {
	var book_type entity.BookType
	if err := c.ShouldBindJSON(&book_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&book_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book_type})
}

// GET /resolution/:id
func GetBookType(c *gin.Context) {
	var book_type entity.BookType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM book_types WHERE id = ?", id).Scan(&book_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book_type})
}

// GET /resolutions
func ListBookType(c *gin.Context) {
	var book_types []entity.BookType
	if err := entity.DB().Raw("SELECT * FROM book_types").Scan(&book_types).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book_types})
}

// DELETE /resolutions/:id
func DeleteBookType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM book_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book_type not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /resolutions
func UpdateBookType(c *gin.Context) {
	var book_type entity.BookType
	if err := c.ShouldBindJSON(&book_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", book_type.ID).First(&book_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booktypes not found"})
		return
	}

	if err := entity.DB().Save(&book_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book_type})
}

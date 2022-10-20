package controller

import (
	"net/http"

	"github.com/B6238285/week5/entity"
	"github.com/gin-gonic/gin"
)

// POST /resolutions
func CreateShelf(c *gin.Context) {
	var shelf entity.Shelf
	if err := c.ShouldBindJSON(&shelf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&shelf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": shelf})
}

// GET /resolution/:id
func GetShelf(c *gin.Context) {
	var shelf entity.Shelf
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM shelves WHERE id = ?", id).Scan(&shelf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shelf})
}

// GET /resolutions
func ListShelf(c *gin.Context) {
	var shelves []entity.Shelf
	if err := entity.DB().Raw("SELECT * FROM shelves").Scan(&shelves).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shelves})
}

// DELETE /resolutions/:id
func DeleteShelf(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM shelves WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shelf not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /resolutions
func UpdateShelf(c *gin.Context) {
	var shelf entity.Shelf
	if err := c.ShouldBindJSON(&shelf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", shelf.ID).First(&shelf); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shelf not found"})
		return
	}

	if err := entity.DB().Save(&shelf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shelf})
}

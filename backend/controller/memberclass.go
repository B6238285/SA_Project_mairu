package controller

import (
	"net/http"

	"github.com/B6238285/week5/entity"
	"github.com/gin-gonic/gin"
)

// POST /memberclasses
func CreateMemberClass(c *gin.Context) {
	var memberclass entity.MemberClass
	if err := c.ShouldBindJSON(&memberclass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&memberclass).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": memberclass})
}

// GET /memberclass/:id
func GetMemberClass(c *gin.Context) {
	var memberclass entity.MemberClass
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&memberclass); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MemberClass not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memberclass})
}

// GET /memberclasses
func ListMemberClass(c *gin.Context) {
	var memberclass []entity.MemberClass
	if err := entity.DB().Raw("SELECT * FROM member_classes").Scan(&memberclass).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memberclass})
}

// DELETE /memberclasses/:id
func DeleteMemberClass(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM member_classes WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MemberClass not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /memberclasses
func UpdateMemberclass(c *gin.Context) {
	var memberclass entity.MemberClass
	if err := c.ShouldBindJSON(&memberclass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", memberclass.ID).First(&memberclass); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MemberClass not found"})
		return
	}

	if err := entity.DB().Save(&memberclass).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memberclass})
}

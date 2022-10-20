package controller

import (
	"net/http"

	"github.com/B6238285/week5/entity"
	"github.com/gin-gonic/gin"
)

// GET /Employee/:id
func GetEmployee(c *gin.Context) {
	var employee entity.Employee
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// GET /employees
func ListEmployee(c *gin.Context) {
	var employee []entity.Employee
	if err := entity.DB().Raw("SELECT * FROM Employees").Scan(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// DELETE /employees/:id
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Employees WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /employees
func UpdateEmployee(c *gin.Context) {
	var employee entity.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", employee.ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found"})
		return
	}

	if err := entity.DB().Save(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

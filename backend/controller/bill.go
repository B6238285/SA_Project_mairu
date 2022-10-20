package controller

import (
	"net/http"

	"github.com/B6238285/week5/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /bills
func CreateBill(c *gin.Context) {

	var bill entity.Bill
	var book entity.Book
	//var memberclass entity.MemberClass
	var user entity.User
	var employee entity.Employee

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา book ด้วย id
	if tx := entity.DB().Where("id = ?", bill.BookID).First(&book); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not  found"})
		return
	}

	// ค้นหา memberclass ด้วย id
	//if tx := entity.DB().Where("id = ?", bill.MemberClass_ID).First(&memberclass); tx.RowsAffected == 0 {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "memberclass not  found"})
	//	return
	//}

	// ค้นหา employe ด้วย id
	if tx := entity.DB().Where("id = ?", bill.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not  found"})
		return
	}

	// ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", bill.UserID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not  found"})
		return
	}

	// สร้าง bill
	bl := entity.Bill{
		Book:     book,          // โยงความสัมพันธ์กับ Entity Book
		Employee: employee,      // โยงความสัมพันธ์กับ Entity Employee
		User:     user,          // โยงความสัมพันธ์กับ Entity User
		BillTime: bill.BillTime, // ตั้งค่าฟิลด์ billTime
	}

	// ขั้นตอน validate
	if _, err := govalidator.ValidateStruct(bl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// บันทึก
	if err := entity.DB().Create(&bl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bl})

}

// GET /bill /:id
func GetBill(c *gin.Context) {
	var bill entity.Bill
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&bill); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bill not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bill})

}

// GET /bills

func ListBills(c *gin.Context) {
	var bills []entity.Bill
	if err := entity.DB().Preload("Book").Preload("Employee").Preload("User").Raw("SELECT * FROM bills").Find(&bills).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bills})
}

// DELETE /bills /:id
func DeleteBill(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM bills WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bill not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /bills
func UpdateBill(c *gin.Context) {
	var bill entity.Bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bill.ID).First(&bill); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bill not found"})
		return
	}

	if err := entity.DB().Save(&bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bill})
}

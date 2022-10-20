package entity

import (
	"time"

	"gorm.io/gorm"
)

// ///////////////////////////////////////////////////////////////////////////
type Employee struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string

	Users []User `gorm:"foreignKey:EmployeeID"`
	Books []Book `gorm:"foreignKey:EmployeeID"`
	Bills []Bill `gorm:"foreignKey:EmployeeID"`
}

type Role struct {
	gorm.Model
	Name       string
	BorrowDay  int
	BookRoomHR int
	BookComHR  int
	Users      []User `gorm:"foreignKey:RoleID"`
	Books      []Book `gorm:"foreignKey:RoleID"`
}

type Province struct {
	gorm.Model
	Name  string
	Users []User `gorm:"foreignKey:ProvinceID"`
}

type MemberClass struct {
	gorm.Model
	Name     string
	Discount int
	Users    []User `gorm:"foreignKey:MemberClassID"`
}

type User struct {
	gorm.Model
	Pin       string `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	Civ       string `gorm:"uniqueIndex"`
	Phone     string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Address   string
	//FK
	EmployeeID    *uint
	RoleID        *uint
	ProvinceID    *uint
	MemberClassID *uint
	//JOIN
	Province    Province    `gorm:"references:id"`
	Role        Role        `gorm:"references:id"`
	MemberClass MemberClass `gorm:"references:id"`
	Employee    Employee    `gorm:"references:id"`
	Bills       []Bill      `gorm:"foreignKey:UserID"`
}

// ///////////////////////////////////////////////////////////////////////////
type BookType struct {
	gorm.Model
	Type string
	//1 book type มีได้หลาย book
	Books []Book `gorm:"foreignKey:BooktypeID"`
}

type Shelf struct {
	gorm.Model
	Type  string
	Floor uint
	//1 shelf มีได้หลาย book
	Books []Book `gorm:"foreignKey:ShelfID"`
}

type Book struct {
	gorm.Model
	Name string
	//ทำหน้าที่เป็น FK
	EmployeeID *uint
	BooktypeID *uint
	ShelfID    *uint
	RoleID     *uint
	//join ให้งายขึ้น
	Employee Employee `gorm:"references:id"`
	Booktype BookType `gorm:"references:id"`
	Shelf    Shelf    `gorm:"references:id"`
	Role     Role     `gorm:"references:id"`
	Author   string
	Page     int
	Quantity int
	Price    int
	Date     time.Time
	Bills    []Bill `gorm:"foreignKey:BookID"`
}

/////////////////////////////////////////////////////////////////////////////

type Bill struct { //เป็นการ get api มาจาก code จะไปอยู่ในส่วนของ front end
	gorm.Model
	//ทำหน้าที่เป็น FK
	BookID *uint
	Book   Book `gorm:"references:id"`

	//ทำหน้าที่เป็น FK
	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	//ทำหน้าที่เป็น FK
	//MemberClass_ID *uint

	//ทำหน้าที่เป็น FK
	UserID *uint
	User   User `gorm:"references:id"`

	//join ให้งายขึ้น

	Book_Name        string
	MemberClass_Name string
	Book_Price       uint //uint ไม่มีเครื่องหมายติดลบ
	Discount         uint
	Total            uint
	BillTime         time.Time
}

package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&Employee{},
		&Role{},
		&Province{},
		&MemberClass{},
		&User{},
		&BookType{},
		&Shelf{},
		&Book{},
		&Bill{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&Employee{}).Create(&Employee{
		Name:     "Sirinya",
		Email:    "sirinya@mail.com",
		Password: "zaq1@wsX",
	})

	db.Model(&Employee{}).Create(&Employee{
		Name:     "Attawit",
		Email:    "attawit@mail.com",
		Password: "zxvsetabb",
	})

	var sirin Employee
	var attawit Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "sirinya@mail.com").Scan(&sirin)
	db.Raw("SELECT * FROM employees WHERE email = ?", "attawit@mail.com").Scan(&attawit)

	//Role
	student := Role{
		Name:       "Student",
		BorrowDay:  3,
		BookRoomHR: 3,
		BookComHR:  4,
	}

	db.Model(&Role{}).Create(&student)

	teacher := Role{
		Name:       "Teacher",
		BorrowDay:  7,
		BookRoomHR: 12,
		BookComHR:  12,
	}
	db.Model(&Role{}).Create(&teacher)

	//province
	korat := Province{
		Name: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&korat)

	chon := Province{
		Name: "Chonburi",
	}
	db.Model(&Province{}).Create(&chon)

	bangkok := Province{
		Name: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	//member
	classic := MemberClass{
		Name:     "classic",
		Discount: 0,
	}
	db.Model(&MemberClass{}).Create(&classic)

	silver := MemberClass{
		Name:     "silver",
		Discount: 5,
	}
	db.Model(&MemberClass{}).Create(&silver)

	gold := MemberClass{
		Name:     "gold",
		Discount: 10,
	}
	db.Model(&MemberClass{}).Create(&gold)

	plat := MemberClass{
		Name:     "platinum",
		Discount: 20,
	}
	db.Model(&MemberClass{}).Create(&plat)

	//user
	db.Model(&User{}).Create(&User{
		Pin:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		Civ:       "1111111111111",
		Phone:     "0811111111",
		Email:     "preechapat@mail.com",
		Password:  string(password),
		Address:   "ถนน a อำเภอ v",
		//FK
		Employee:    sirin,
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	})

	//Shelf
	S1 := Shelf{
		Type:  "SCIENCE",
		Floor: 1,
	}
	db.Model(&Shelf{}).Create(&S1)
	S2 := Shelf{
		Type:  "ENGINEERING",
		Floor: 1,
	}
	db.Model(&Shelf{}).Create(&S2)
	S3 := Shelf{
		Type:  "ENVIRRONMENT",
		Floor: 1,
	}
	db.Model(&Shelf{}).Create(&S3)
	S4 := Shelf{
		Type:  "HISTORY",
		Floor: 1,
	}
	db.Model(&Shelf{}).Create(&S4)
	S5 := Shelf{
		Type:  "FICTION",
		Floor: 2,
	}
	db.Model(&Shelf{}).Create(&S5)
	S6 := Shelf{
		Type:  "FANTASY",
		Floor: 2,
	}
	db.Model(&Shelf{}).Create(&S6)
	S7 := Shelf{
		Type:  "HORROR",
		Floor: 2,
	}
	db.Model(&Shelf{}).Create(&S7)

	//Book Type
	BT1 := BookType{
		Type: "COMPUTER ENGINEERING",
	}
	db.Model(&BookType{}).Create(&BT1)

	BT2 := BookType{
		Type: "ELECTRIC ENGINEERING",
	}
	db.Model(&BookType{}).Create(&BT2)

	BT3 := BookType{
		Type: "SUPERHERO FANTASY",
	}
	db.Model(&BookType{}).Create(&BT3)

	BT4 := BookType{
		Type: "HORROR FICTION",
	}
	db.Model(&BookType{}).Create(&BT4)

	BT5 := BookType{
		Type: "DARK AND GRIMDARK FANTASY",
	}
	db.Model(&BookType{}).Create(&BT5)
	BT6 := BookType{
		Type: "CONTEMPORARY FANTASY",
	}
	db.Model(&BookType{}).Create(&BT6)

	//Book
	db.Model(&Book{}).Create(&Book{
		Name:     "Python 1",
		Employee: sirin,
		Booktype: BT1,
		Shelf:    S2,
		Role:     student,
		Author:   "Sirin",
		Page:     500,
		Quantity: 20,
		Price:    300,
		Date:     time.Now(),
	})
	db.Model(&Book{}).Create(&Book{
		Name:     "Java",
		Employee: attawit,
		Booktype: BT1,
		Shelf:    S2,
		Role:     teacher,
		Author:   "AJ",
		Page:     350,
		Quantity: 10,
		Price:    200,
		Date:     time.Now(),
	})
	var Python Book
	db.Raw("SELECT * FROM Books WHERE name = ? ", "Python 1").Scan(&Python) //ดึง id

	var User1 User
	db.Raw("SELECT * FROM Users WHERE pin = ? ", "B6111111").Scan(&User1)      //ดึง id
	db.Raw("SELECT * FROM Users WHERE civ = ? ", "1111111111111").Scan(&User1) //ดึง id

	db.Model(&Bill{}).Create(&Bill{
		Book_Name:        Python.Name, //ค้นหาจาก id
		Book_Price:       uint(Python.Price),
		Employee:         sirin, //ค้นหาจาก id
		Book:             Python,
		User:             User1,                  //ค้นหาจาก id
		MemberClass_Name: User1.MemberClass.Name, //ดึงไงวะ
		Discount:         20,
		Total:            480,
		BillTime:         time.Now(),
	})

}

package main

import (
	"github.com/B6238285/week5/controller"
	"github.com/B6238285/week5/entity"

	//"github.com/B6238285/week5/middlewares"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	//r.Use(CORSMiddleware())

	router := r.Group("/")

	{
		//router.Use(middlewares.Authorizes())
		{
			router.GET("/users", controller.ListUser)
			router.GET("/user/:id", controller.GetUser)
			router.POST("/users", controller.CreateUser)
			router.PATCH("/users", controller.UpdateUser)
			router.DELETE("/users/:id", controller.DeleteUser)

			//employee routes
			router.GET("/employees", controller.ListEmployee)
			router.GET("/employee/:id", controller.GetEmployee)
			//r.POST("/employees", controller.CreateEmployee)
			router.PATCH("/employees", controller.UpdateEmployee)
			router.DELETE("/employees/:id", controller.DeleteEmployee)

			//memberClass routes
			router.GET("/memberclasses", controller.ListMemberClass)
			router.GET("/memberclass/:id", controller.GetMemberClass)
			router.POST("/memberclasses", controller.CreateMemberClass)
			router.PATCH("/memberclasses", controller.UpdateMemberclass)
			router.DELETE("/memberclasses/:id", controller.DeleteMemberClass)

			//province routes
			router.GET("/provinces", controller.ListProvince)
			router.GET("/province/:id", controller.GetProvince)
			router.POST("/provinces", controller.CreateProvince)
			router.PATCH("/provinces", controller.UpdateProvince)
			router.DELETE("/provinces/:id", controller.DeleteMemberClass)

			//role routes
			router.GET("/roles", controller.ListRole)
			router.GET("/role/:id", controller.GetRole)
			router.POST("/roles", controller.CreateUser)
			router.PATCH("/roles", controller.UpdateRole)
			router.DELETE("/roles/:id", controller.DeleteRole)

			// shelf Routes
			router.GET("/shelves", controller.ListShelf)
			router.GET("/shelve/:id", controller.GetShelf)
			router.POST("/shelves", controller.CreateShelf)
			router.PATCH("/shelves", controller.UpdateShelf)
			router.DELETE("/shelves/:id", controller.DeleteShelf)

			// BOOK_TYPE Routes
			router.GET("/book_types", controller.ListBookType)
			router.GET("/book_types/:id", controller.GetBookType)
			router.POST("/book_types", controller.CreateBookType)
			router.PATCH("/book_types", controller.UpdateBookType)
			router.DELETE("/book_types/:id", controller.DeleteBookType)

			// book Routes
			router.GET("/books", controller.ListBook)
			router.GET("/book/:id", controller.GetBook)
			router.POST("/books", controller.CreateBook)
			router.PATCH("/books", controller.UpdateBook)
			router.DELETE("/books/:id", controller.DeleteBook)

			//bill route
			router.GET("/bills", controller.ListBills)
			router.GET("/bill/:id", controller.GetBill)
			router.POST("/bills", controller.CreateBill)
			router.PATCH("/bills", controller.UpdateBill)
			router.DELETE("/bills/:id", controller.DeleteBill)

			// Run the server
			r.Run()
		}
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

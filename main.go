package main

import (
	"code/database/config"
	"code/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	SetupDB()
	SetupRoutes()
}

func SetupDB() {
	config.Setup()
}

func SetupRoutes() {
	route := gin.Default()
	route.GET("/api/books", routes.GetBooks)
	route.GET("/api/books/find/by/book/:id", routes.GetBook)
	route.GET("/api/books/find/by/category/:id", routes.GetBookByCategory)
	route.POST("/api/books/create", routes.CreateBook)
	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

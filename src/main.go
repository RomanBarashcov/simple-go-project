package main

import (
	"simple-go-project/src/database/config"
	"simple-go-project/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	SetupDB()
	MakeInjection()
	SetupRoutes()

}

func SetupDB() {

	config.Setup()

}

func SetupRoutes() {

	route := gin.Default()

	SetupBookRoutes(route)
	SetupReviewRoutes(route)

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func SetupBookRoutes(route *gin.Engine) {

	route.GET("/api/books", routes.GetBooks)
	route.GET("/api/books/find/by/book/:id", routes.GetBook)
	route.GET("/api/books/find/by/category/:id", routes.GetBookByCategory)
	route.POST("/api/books/create", routes.CreateBook)
	route.PUT("/api/books/update", routes.UpdateBook)
	route.DELETE("/api/books/delete/:id", routes.DeleteBook)

}

func SetupReviewRoutes(route *gin.Engine) {

	route.POST("/api/reviews/create", routes.CreateReview)
	route.PUT("/api/reviews/update", routes.UpdateReview)
	route.DELETE("/api/reviews/delete/:id", routes.DeleteReview)

}

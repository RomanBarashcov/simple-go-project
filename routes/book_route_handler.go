package routes

import (
	"code/entities"
	"code/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {

	books := services.GetAllBooks()
	c.JSON(200, books)

}

func GetBook(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, "Payload is incorrect!")
	}

	book := services.GetBookById(id)
	c.JSON(200, book)
}

func GetBookByCategory(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, "Payload is incorrect!")
	}

	book := services.GetBooksByCategoryId(id)
	c.JSON(200, book)

}

func CreateBook(c *gin.Context) {

	id, err := strconv.ParseInt(c.PostForm("CategoryID"), 10, 64)
	if err != nil {
		c.JSON(500, "Payload is incorrect!")
	}

	price, err := strconv.ParseInt(c.PostForm("Price"), 10, 32)
	if err != nil {
		c.JSON(500, "Payload is incorrect!")
	}

	newBook := new(entities.Book)

	newBook.Title = c.PostForm("Title")
	newBook.Description = c.PostForm("Description")
	newBook.Cover = c.PostForm("Cover")
	newBook.Price = int32(price)
	newBook.CategoryID = id

	book := services.CreateBook(newBook)

	c.JSON(200, book)

}

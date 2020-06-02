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

	newBook := &entities.Book{}
	c.Bind(newBook)

	book := services.CreateBook(newBook)

	c.JSON(200, book)

}

func UpdateBook(c *gin.Context) {

	upBook := &entities.Book{}
	c.Bind(upBook)

	book := services.UpdateBook(upBook)

	c.JSON(200, book)

}

func DeleteBook(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, "Payload is incorrect!")
	}

	success := services.DeleteBook(id)

	c.PureJSON(200, gin.H{
		"deleted": success,
	})
}

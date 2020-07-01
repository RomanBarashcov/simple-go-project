package tests

import (
	"simple-go-project/src/entities"
	utils "simple-go-project/src/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func StartBookServiceTesting() {
	utils.WriteConsoleLog("Start BookService testring")
}

func Test_ServiceFindAllBooks(t *testing.T) {

	StartBookServiceTesting()
	SetupServiceIngection()

	actual := bSer.GetAllBooks()
	expected := bRep.FindAllBooks()

	assert.Equal(t, expected[0].ID, actual[0].ID)

}

func Test_ServiceGetBookById(t *testing.T) {

	bookID := int64(1)
	book, _ := bSer.GetBookById(bookID)
	assert.Equal(t, book.ID, bookID)

}

func Test_ByIncorrectID_ServiceGetBookById(t *testing.T) {

	bookID := int64(2)
	_, err := bSer.GetBookById(bookID)
	assert.EqualError(t, err, "Book does not exist")

}

func Test_ServiceCreateBook(t *testing.T) {

	book := new(entities.Book)

	book.Title = "Title4"
	book.Description = "Description4"
	book.Author = "Author4"
	book.Cover = "Cover4"
	book.Price = 400
	book.CategoryName = "CategoryName4"
	book.CategoryID = 3

	actual, _ := bSer.CreateBook(book)
	book.ID = 5
	expected := book

	assert.Equal(t, actual.ID, expected.ID)

}

func Test_IncorrectCategoryID_ServiceCreateBook(t *testing.T) {

	book := new(entities.Book)

	book.Title = "Title4"
	book.Description = "Description4"
	book.Author = "Author4"
	book.Cover = "Cover4"
	book.Price = 400
	book.CategoryName = "CategoryName4"
	book.CategoryID = 30

	_, err := bSer.CreateBook(book)

	assert.EqualError(t, err, "Category does not exist")

}

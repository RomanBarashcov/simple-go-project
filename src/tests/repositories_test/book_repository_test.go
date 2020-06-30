package repositories_test

import (
	"simple-go-project/src/entities"
	utils "simple-go-project/src/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func StartBookRepositoryTesting() {
	utils.WriteConsoleLog("Start BookRepository testring")
}

func TestFindAllBooks(t *testing.T) {

	StartBookRepositoryTesting()

	repository := &MBookRepository{}
	actual := repository.FindAllBooks()
	expected := mockBooks

	assert.Equal(t, expected[0].ID, actual[0].ID)

}

func TestFindBookById(t *testing.T) {

	repository := &MBookRepository{}
	var bookId int64 = 1

	actual := repository.FindBookById(bookId)
	expected := mockBooks[0].ID

	assert.Equal(t, expected, actual.ID)

}

func TestFindBooksByCategoryId(t *testing.T) {

	repository := &MBookRepository{}
	var categoryId int64 = 1

	actual := repository.FindBooksByCategoryId(categoryId)
	expected := mockBooks

	assert.Equal(t, expected[0].ID, actual[0].ID)

}

func TestCreateBook(t *testing.T) {

	repository := &MBookRepository{}
	book := new(entities.Book)

	book.Title = "Title4"
	book.Description = "Description4"
	book.Author = "Author4"
	book.Cover = "Cover4"
	book.Price = 400
	book.CategoryName = "CategoryName4"
	book.CategoryID = 4

	actual := repository.CreateBook(book)
	book.ID = 4
	expected := book

	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, mockBooks[3].ID, actual.ID)
}

func TestUpdateBook(t *testing.T) {

	repository := &MBookRepository{}
	upBook := mockBooks[1]

	upBook.Title = "Title5"
	upBook.Description = "Description5"
	upBook.Author = "Author5"
	upBook.Cover = "Cover5"
	upBook.Price = 500
	upBook.CategoryName = "CategoryName5"
	upBook.CategoryID = 5

	_ = repository.UpdateBook(&upBook)
	expected := upBook

	assert.Equal(t, expected.ID, mockBooks[1].ID)
	assert.Equal(t, expected.Title, mockBooks[1].Title)

}

func TestDeleteBook(t *testing.T) {

	repository := &MBookRepository{}
	delBook := mockBooks[1]

	actual := repository.DeleteBook(delBook.ID)
	expected := true

	assert.Equal(t, actual, expected)
	assert.NotEqual(t, delBook.ID, mockBooks[1].ID)

	EndBookRepositoryTesting()

}

func EndBookRepositoryTesting() {
	utils.WriteConsoleLog("End BookRepository testring")
}

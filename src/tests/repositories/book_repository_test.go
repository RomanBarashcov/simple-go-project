package repositories_test

import (
	"simple-go-project/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllBooks(t *testing.T) {

	repository := &MockBookRepository{}
	actual := repository.FindAllBooks()
	expected := mockBooks

	assert.Equal(t, expected[0].ID, actual[0].ID)

}

func TestFindBookById(t *testing.T) {

	repository := &MockBookRepository{}
	var bookId int64 = 1

	actual := repository.FindBookById(bookId)
	expected := mockBooks[0].ID

	assert.Equal(t, expected, actual.ID)

}

func TestFindBooksByCategoryId(t *testing.T) {

	repository := &MockBookRepository{}
	var categoryId int64 = 1

	actual := repository.FindBooksByCategoryId(categoryId)
	expected := mockBooks

	assert.Equal(t, expected[0].ID, actual[0].ID)

}

func TestCreateBook(t *testing.T) {

	repository := &MockBookRepository{}
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

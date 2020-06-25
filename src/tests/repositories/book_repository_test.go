package repositories_test

import (
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

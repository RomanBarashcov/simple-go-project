package services

import (
	"fmt"
	"simple-go-project/src/entities"
)

type IBookService interface {
	GetAllBooks() []entities.Book
	GetBookById(id int64) (*entities.Book, error)
	GetBooksByCategoryId(id int64) ([]entities.Book, error)
	CreateBook(newBook *entities.Book) (*entities.Book, error)
	UpdateBook(upBook *entities.Book) (*entities.Book, error)
	DeleteBook(id int64) bool
}

type BookService struct{}

func (b *BookService) GetAllBooks() []entities.Book {
	books := bookRepository.FindAllBooks()
	return books
}

func (b *BookService) GetBookById(id int64) (*entities.Book, error) {

	book := bookRepository.FindBookById(id)
	if book == nil {
		return nil, fmt.Errorf("Book does not exist")
	}

	return book, nil
}

func (b *BookService) GetBooksByCategoryId(id int64) ([]entities.Book, error) {

	books := bookRepository.FindBooksByCategoryId(id)
	if books == nil {
		return nil, fmt.Errorf("Category does not exist")
	}

	return books, nil
}

func (b *BookService) CreateBook(newBook *entities.Book) (*entities.Book, error) {

	category := bookRepository.FindCategoryById(newBook.CategoryID)
	if category == nil {
		return nil, fmt.Errorf("Category does not exist")
	}

	book := bookRepository.CreateBook(newBook)

	return book, nil
}

func (b *BookService) UpdateBook(upBook *entities.Book) (*entities.Book, error) {

	category := bookRepository.FindBooksByCategoryId(upBook.CategoryID)
	if category == nil {
		return nil, fmt.Errorf("Category does not exist")
	}

	book := bookRepository.UpdateBook(upBook)

	return book, nil
}

func (b *BookService) DeleteBook(id int64) bool {
	success := bookRepository.DeleteBook(id)
	return success
}

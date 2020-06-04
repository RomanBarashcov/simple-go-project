package services

import (
	"code/entities"
)

type IBookService interface {
	GetAllBooks() []entities.Book
	GetBookById(id int64) *entities.Book
	GetBooksByCategoryId(id int64) []entities.Book
	CreateBook(newBook *entities.Book) *entities.Book
	UpdateBook(upBook *entities.Book) *entities.Book
	DeleteBook(id int64) bool
}

type BookService struct{}

func (b *BookService) GetAllBooks() []entities.Book {
	books := bookRepository.FindAllBooks()
	return books
}

func (b *BookService) GetBookById(id int64) *entities.Book {
	book := bookRepository.FindBookById(id)
	return book
}

func (b *BookService) GetBooksByCategoryId(id int64) []entities.Book {
	books := bookRepository.FindBooksByCategoryId(id)
	return books
}

func (b *BookService) CreateBook(newBook *entities.Book) *entities.Book {
	book := bookRepository.CreateBook(newBook)
	return book
}

func (b *BookService) UpdateBook(upBook *entities.Book) *entities.Book {
	book := bookRepository.UpdateBook(upBook)
	return book
}

func (b *BookService) DeleteBook(id int64) bool {
	success := bookRepository.DeleteBook(id)
	return success
}

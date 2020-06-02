package services

import (
	"code/entities"
	"code/repositories"
)

func GetAllBooks() []entities.Book {
	books := repositories.FindAllBooks()
	return books
}

func GetBookById(id int64) *entities.Book {
	book := repositories.FindBookById(id)
	return book
}

func GetBooksByCategoryId(id int64) []entities.Book {
	books := repositories.FindBooksByCategoryId(id)
	return books
}

func CreateBook(newBook *entities.Book) *entities.Book {
	book := repositories.CreateBook(newBook)
	return book
}

func UpdateBook(upBook *entities.Book) *entities.Book {
	book := repositories.UpdateBook(upBook)
	return book
}

func DeleteBook(id int64) bool {
	success := repositories.DeleteBook(id)
	return success
}

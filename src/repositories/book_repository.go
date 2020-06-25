package repositories

import (
	"simple-go-project/src/database/config"
	"simple-go-project/src/database/models"
	"simple-go-project/src/entities"
)

type IBookRepository interface {
	FindAllBooks() []entities.Book
	FindBookById(id int64) *entities.Book
	FindBooksByCategoryId(id int64) []entities.Book
	CreateBook(newBook *entities.Book) *entities.Book
	UpdateBook(upBook *entities.Book) *entities.Book
	DeleteBook(id int64) bool
}

type BookRepository struct{}

func (b *BookRepository) FindAllBooks() []entities.Book {

	db := config.GetConnection()

	rows, err := db.Table("books as b").Select("b.*, c.name as category_name").Joins("left join categories as c on c.ID = b.category_id").Rows()
	if err != nil {
		panic("Some error")
	}

	books := make([]entities.Book, 0)

	for rows.Next() {
		var book entities.Book
		db.ScanRows(rows, &book)
		books = append(books, book)
	}

	return books

}

func (b *BookRepository) FindBookById(id int64) *entities.Book {

	db := config.GetConnection()

	book := new(entities.Book)
	db.Table("books as b").Select("b.*, c.name as category_name").Where("b.ID = ?", id).Joins("left join categories as c on c.ID = b.category_id").Scan(&book)
	db.Table("reviews as r").Select("r.*, u.name as user_name, u.email as user_email").Where("r.book_id = ?", id).Joins("left join users as u on u.id = r.user_id").Scan(&book.Reviews)

	return book

}

func (b *BookRepository) FindBooksByCategoryId(id int64) []entities.Book {

	db := config.GetConnection()

	rows, err := db.Table("books as b").Select("b.*, c.name as category_name").Where("b.category_id = ?", id).Joins("left join categories as c on c.ID = b.category_id").Rows()
	if err != nil {
		panic("Some error")
	}

	books := make([]entities.Book, 0)

	for rows.Next() {
		var book entities.Book
		db.ScanRows(rows, &book)
		books = append(books, book)
	}

	return books

}

func (b *BookRepository) CreateBook(newBook *entities.Book) *entities.Book {

	db := config.GetConnection()

	book := models.Book{
		Title:       newBook.Title,
		Description: newBook.Description,
		Author:      newBook.Author,
		Cover:       newBook.Cover,
		Price:       newBook.Price,
		CategoryID:  newBook.CategoryID,
	}

	newRecordResult := db.NewRecord(book)
	if !newRecordResult {
		panic("new Record book error")
	}

	db.Create(&book).Scan(&newBook)
	return newBook
}

func (b *BookRepository) UpdateBook(upBook *entities.Book) *entities.Book {

	db := config.GetConnection()

	book := models.Book{ID: upBook.ID}
	db.First(&book)

	book.Title = upBook.Title
	book.Description = upBook.Description
	book.Author = upBook.Author
	book.Cover = upBook.Cover
	book.Price = upBook.Price
	book.CategoryID = upBook.CategoryID

	upResult := db.Save(&book).Value
	if upResult == nil {
		panic("Error when update book")
	}

	return upBook

}

func (b *BookRepository) DeleteBook(id int64) bool {

	var success bool = false

	db := config.GetConnection()

	book := models.Book{ID: id}
	result := db.Delete(&book).Value
	if result == nil {
		panic("Error when delete book")
	}

	success = true
	return success
}

package repositories

import (
	"code/database/config"
	"code/database/models"
	"code/entities"
)

func FindAllBooks() []entities.Book {

	db := config.GetConnection()
	defer db.Close()

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

func FindBookById(id int64) *entities.Book {

	db := config.GetConnection()
	defer db.Close()

	book := new(entities.Book)
	db.Table("books as b").Select("b.*, c.name as category_name").Where("b.ID = ?", id).Joins("left join categories as c on c.ID = b.category_id").Scan(&book)

	return book

}

func FindBooksByCategoryId(id int64) []entities.Book {

	db := config.GetConnection()
	defer db.Close()

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

func CreateBook(newBook *entities.Book) *entities.Book {

	db := config.GetConnection()
	defer db.Close()

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

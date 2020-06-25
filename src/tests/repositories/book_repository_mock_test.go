package repositories_test

import (
	"simple-go-project/src/entities"
)

type MockBookRepository struct {
}

var mockBooks = []entities.Book{

	entities.Book{
		ID:           1,
		Title:        "Title1",
		Description:  "Description1",
		Author:       "Author1",
		Cover:        "Cover1",
		Price:        100,
		CategoryName: "Arts & Photography",
		CategoryID:   1,
		Reviews:      nil,
	},
	entities.Book{
		ID:           2,
		Title:        "Title2",
		Description:  "Description2",
		Author:       "Author2",
		Cover:        "Cover2",
		Price:        200,
		CategoryName: "Arts & Photography",
		CategoryID:   2,
		Reviews:      nil,
	},
	entities.Book{
		ID:           3,
		Title:        "Title3",
		Description:  "Description3",
		Author:       "Author3",
		Cover:        "Cover3",
		Price:        300,
		CategoryName: "Arts & Photography",
		CategoryID:   3,
		Reviews:      nil,
	},
}

func (m *MockBookRepository) FindAllBooks() []entities.Book {

	books := mockBooks
	return books

}

func (m *MockBookRepository) FindBookById(id int64) *entities.Book {

	book := new(entities.Book)

	for _, v := range mockBooks {

		if v.ID == id {

			book.ID = v.ID
			book.Title = v.Title
			book.Description = v.Description
			book.Author = v.Author
			book.Cover = v.Cover
			book.Price = v.Price
			book.CategoryName = v.CategoryName
			book.CategoryID = v.CategoryID
			book.Reviews = v.Reviews

			break

		}

	}

	return book
}

func (m *MockBookRepository) FindBooksByCategoryId(id int64) []entities.Book {

	books := make([]entities.Book, 0)

	for _, v := range mockBooks {

		if v.CategoryID == id {

			books = append(books, v)
		}

	}

	return books
}

func (m *MockBookRepository) CreateBook(newBook *entities.Book) *entities.Book {

	return nil
}

func (m *MockBookRepository) UpdateBook(upBook *entities.Book) *entities.Book {
	return nil
}

func (m *MockBookRepository) DeleteBook(id int64) bool {
	return true
}

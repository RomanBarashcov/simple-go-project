package repositories_test

import (
	"simple-go-project/src/entities"
	"simple-go-project/src/repositories"
)

type mockBookRepository struct{}

var (
	MockBookRepository repositories.IBookRepository = &mockBookRepository{}
)

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

func (m *mockBookRepository) FindAllBooks() []entities.Book {

	books := mockBooks
	return books

}

func (m *mockBookRepository) FindBookById(id int64) *entities.Book {

	for i, v := range mockBooks {

		if v.ID == id {

			return &mockBooks[i]

		}

	}

	return nil
}

func (m *mockBookRepository) FindBooksByCategoryId(id int64) []entities.Book {

	books := make([]entities.Book, 0)

	for _, v := range mockBooks {

		if v.CategoryID == id {

			books = append(books, v)
		}

	}

	return books
}

func (m *mockBookRepository) CreateBook(newBook *entities.Book) *entities.Book {

	newBook.ID = int64(len(mockBooks)) + 1
	mockBooks = append(mockBooks, *newBook)
	return newBook

}

func (m *mockBookRepository) UpdateBook(upBook *entities.Book) *entities.Book {

	for i, v := range mockBooks {

		if v.ID == upBook.ID {

			mockBooks[i].Title = upBook.Title
			mockBooks[i].Description = upBook.Description
			mockBooks[i].Author = upBook.Author
			mockBooks[i].Cover = upBook.Cover
			mockBooks[i].Price = upBook.Price
			mockBooks[i].CategoryName = upBook.CategoryName
			mockBooks[i].CategoryID = upBook.CategoryID
			mockBooks[i].Reviews = upBook.Reviews

		}

	}

	return upBook

}

func (m *mockBookRepository) DeleteBook(id int64) bool {

	for i, v := range mockBooks {

		if v.ID == id {

			mockBooks = append(mockBooks[:i], mockBooks[i+1:]...) // delete element by index
			return true
		}

	}

	return false
}

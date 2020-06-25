package config

import (
	"simple-go-project/src/database/models"

	"github.com/jinzhu/gorm"
)

func SetupSeed(db *gorm.DB) {

	SetupCategories(db)
	SetupBooks(db)
	SetupUsers(db)
	SetupReviews(db)

}

func SetupCategories(db *gorm.DB) {

	category1 := models.Category{}
	category2 := models.Category{}
	category3 := models.Category{}

	db.Where(models.Category{Name: "Arts & Photography"}).FirstOrCreate(&category1)
	db.Where(models.Category{Name: "Biographies & Memoirs"}).FirstOrCreate(&category2)
	db.Where(models.Category{Name: "Business & Investing"}).FirstOrCreate(&category3)
}

func SetupBooks(db *gorm.DB) {

	book1 := models.Book{}
	book2 := models.Book{}
	book3 := models.Book{}

	db.Where(models.Book{Title: "Title1", Description: "Description1", Cover: "Cover1", Author: "Author1", Price: 100, CategoryID: 1}).FirstOrCreate(&book1)
	db.Where(models.Book{Title: "Title2", Description: "Description2", Cover: "Cover2", Author: "Author2", Price: 200, CategoryID: 2}).FirstOrCreate(&book2)
	db.Where(models.Book{Title: "Title3", Description: "Description3", Cover: "Cover3", Author: "Author3", Price: 300, CategoryID: 3}).FirstOrCreate(&book3)

}

func SetupUsers(db *gorm.DB) {

	user1 := models.User{}
	user2 := models.User{}
	user3 := models.User{}

	db.Where(models.User{Name: "User1", Email: "user1@gmail.com"}).FirstOrCreate(&user1)
	db.Where(models.User{Name: "User2", Email: "user2@gmail.com"}).FirstOrCreate(&user2)
	db.Where(models.User{Name: "User3", Email: "user3@gmail.com"}).FirstOrCreate(&user3)

}

func SetupReviews(db *gorm.DB) {

	review1 := models.Review{}
	review2 := models.Review{}
	review3 := models.Review{}
	review4 := models.Review{}

	db.Where(models.Review{BookID: 1, Rating: 5, Text: "Best book I have ever read1", UserID: 1}).FirstOrCreate(&review1)
	db.Where(models.Review{BookID: 2, Rating: 4, Text: "Best book I have ever read2", UserID: 1}).FirstOrCreate(&review2)
	db.Where(models.Review{BookID: 2, Rating: 3, Text: "Best book I have ever read3", UserID: 2}).FirstOrCreate(&review3)
	db.Where(models.Review{BookID: 2, Rating: 1, Text: "Best book I have ever read4", UserID: 2}).FirstOrCreate(&review4)

}

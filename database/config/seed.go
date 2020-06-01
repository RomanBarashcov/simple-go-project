package config

import "code/database/models"
import "github.com/jinzhu/gorm"

func SetupSeed(db *gorm.DB) {
	SetupCategories(db)
	SetupBooks(db)
}

func SetupCategories(db *gorm.DB) {

	category := models.Category{}

	db.Where(models.Category{Name: "Arts & Photography"}).FirstOrCreate(&category)
	db.Where(models.Category{Name: "Biographies & Memoirs"}).FirstOrCreate(&category)
	db.Where(models.Category{Name: "Business & Investing"}).FirstOrCreate(&category)
}

func SetupBooks(db *gorm.DB) {

	book := models.Book{}

	db.Where(models.Book{Title: "Title1", Description: "Description1", Cover: "Cover1", Author: "Author1", Price: 100, CategoryID: 1}).FirstOrCreate(&book)
	db.Where(models.Book{Title: "Title2", Description: "Description2", Cover: "Cover2", Author: "Author2", Price: 200, CategoryID: 2}).FirstOrCreate(&book)
	db.Where(models.Book{Title: "Title3", Description: "Description3", Cover: "Cover3", Author: "Author3", Price: 300, CategoryID: 3}).FirstOrCreate(&book)

}
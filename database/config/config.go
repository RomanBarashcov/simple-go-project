package config

import (
	m "code/database/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Setup() {

	db := GetConnection()
	defer db.Close()

	SetupMigration(db)
	SetupFK(db)

	SetupSeed(db)
}

func SetupMigration(db *gorm.DB) {
	db.AutoMigrate(&m.Book{}, &m.Category{}, &m.Review{})
}

func SetupFK(db *gorm.DB) {
	db.Model(&m.Book{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	db.Model(&m.Review{}).AddForeignKey("book_id", "books(id)", "RESTRICT", "RESTRICT")
}

func GetConnection() *gorm.DB {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=BookStore password=root")
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

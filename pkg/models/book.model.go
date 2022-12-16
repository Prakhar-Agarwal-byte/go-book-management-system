package models

import (
	"github.com/Prakhar-Agarwal-byte/go-book-management-system/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}       

func CreateBook(book *Book) *Book {
	db.Create(book)
	return book
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", ID).First(&book)
	return &book, db
}

func DeleteBook(ID int64) *Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return &book
}
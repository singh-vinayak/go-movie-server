package models

import (
	"github.com/jinzhu/gorm"
	"github.com/singh-vinayak/go-movie-server/tree/mysql/pkg/config"
)

var db *gorm.DB

type Book struct {
	*gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, * gorm.DB) {
	var book Book
	db := db.Where("ID=?",id).Find(&book)

	return &book, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
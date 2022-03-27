package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pk12799/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBooks() []Book {
	var book []Book
	db.Find(&book)
	return book
}

func GetBook(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("id=?", Id).Find(&getBook)
	log.Println(Id, &getBook, db, "get book model")
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("id=?", ID).Delete(book)
	return book
}

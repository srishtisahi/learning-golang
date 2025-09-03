package models

import (
	"github.com/jinzhu/gorm"
	"github.com/srishtisahi/go-mysql/pkg/config"
)

var db *gorm.db

type Book struct {
	gorm.models // based on structs
	Name string `gorm: ""json: "name"`
	Author string `json: "author"`
	Publication string `json: "publication"`
}

func init() {
	config.Connect()
	db = config.GetDB() // function in app.go
	db.AutoMigrate(&Book{}) // auto migration from config
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // recieve something of type b
	db.Create(&b) // create
	return b
}

func GetAllBooks() []Book {
	var Books []Book // returns a slice (list of books in db)
	db.Find(&Books)
	return Books 
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook) // look for a book and find it
	return &getBook, db // returns book and db variable
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
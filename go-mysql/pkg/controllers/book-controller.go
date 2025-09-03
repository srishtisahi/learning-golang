package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/srishtisahi/go-mysql/pkg/utils" // import utils and models files
	"github.com/srishtisahi/go-mysql/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks() // import func from book.go
	res, _ := json.Marshal(newBooks) // takes db data and marshals it
	w.header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) // gives 200
	w.Write(res) // sends to frontend
}

func GetBookById(w http.ResponseWriter, r http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0) // make sure that bookid is a string
	if err != nil {
		fmt.Println("erorr while parsing")
	}
	bookDetails, _ := models.GetBookById(ID) // use _ whenever you dont use something to avoid errors
	// above line returns book and db from models file
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{} // define a model
	utils.ParseBody(r, CreateBook) // file from utils.go
	b := CreateBook.CreateBook() 
	res, _ := json.Marshal(b) // b was created and returned to us
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.vars(r)
	bookId := vars["bookId"]
	ID, err := strconv
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// same thing as above in getbookbyid
	// access request, bookid, convert to string
	// make sure there was no error
}

func UpdateBook(w http.ResponseWriter, r http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId : vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// the above code is from create and getbookbyid funcs in this file
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name 
	} // make sure the strings are not empty
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author 
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
}


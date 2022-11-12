package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AbraaoLeonardo/golang-projects/pkg/models"
	"github.com/AbraaoLeonardo/golang-projects/pkg/utils"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(200)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.WriteHeader(200)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["BookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Erro while parsing")
	}
	bookDetail, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetail)
	w.WriteHeader(200)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["BookId"]
	book, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Cannot parser")
	}
	deletedBook := models.DeleteBook(book)
	res, _ := json.Marshal(deletedBook)
	w.WriteHeader(200)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["BooBookIdId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Cannot Parse")
	}
	bookDetail, db := models.GetBookById(id)
	if updateBook.Name != "" {
		bookDetail.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetail.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetail.Publication = updateBook.Publication
	}
	db.Save(&bookDetail)
	res, _ := json.Marshal(bookDetail)
	w.WriteHeader(200)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

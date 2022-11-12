package main

// Import the library
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Create a object book
type BookModel struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

// Set the valuer of book
var Books []BookModel = []BookModel{
	{
		Id:     1,
		Name:   "Harry Potter",
		Author: "J.K",
	},
	{
		Id:     2,
		Name:   "Clear Code",
		Author: "Uncle Bob",
	},
	{
		Id:     3,
		Name:   "Frankenstein",
		Author: "Mary Shelley",
	},
}

// Main Route function
func MainRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, 世界")
}

// Create book
func AddNewBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)

	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		fmt.Fprintf(w, "Error: %v", erro)
	}

	var newBook BookModel
	json.Unmarshal(body, &newBook)
	newBook.Id = len(Books) + 1
	Books = append(Books, newBook)

	encoder := json.NewEncoder(w)
	encoder.Encode(newBook)
}

// Read all item
func ShowBooks(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)
	encoder.Encode(Books)

}

// Read 1 item
func SearchBook(w http.ResponseWriter, r *http.Request) {
	SplitedPath := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(SplitedPath[2])

	for _, book := range Books {
		if book.Id == id {
			json.NewEncoder(w).Encode(book)
		}
	}

	w.WriteHeader(404)

}

// Update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	SplitedPath := strings.Split(r.URL.Path, "/")
	id, erro := strconv.Atoi(SplitedPath[2])

	if erro != nil {
		w.WriteHeader(404)
		return
	}

	body, erroBody := io.ReadAll(r.Body)

	if erroBody != nil {
		w.WriteHeader(500)
		return
	}

	var modifiedBook BookModel
	erroJson := json.Unmarshal(body, &modifiedBook)

	if erroJson != nil {
		w.WriteHeader(401)
		return
	}

	var indexBook = -1
	for index, book := range Books {
		if book.Id == id {
			indexBook = index
			break
		}
	}

	if indexBook == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Books[indexBook] = modifiedBook

}

// Delete book
func RemoveBook(w http.ResponseWriter, r *http.Request) {

	SplitedPath := strings.Split(r.URL.Path, "/")
	id, erro := strconv.Atoi(SplitedPath[2])

	if erro != nil {
		w.WriteHeader(404)
		return
	}

	var indexBook = -1
	for index, book := range Books {
		if book.Id == id {
			indexBook = index
			break
		}
	}

	if indexBook == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	left := Books[:indexBook]
	right := Books[indexBook+1:]
	Books = append(left, right...)
	w.WriteHeader(http.StatusNoContent)
}

func RouteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	SplitedPath := strings.Split(r.URL.Path, "/")

	if len(SplitedPath) == 2 || SplitedPath[2] == "" {

		if r.Method == "GET" {
			ShowBooks(w, r)
		} else if r.Method == "POST" {
			AddNewBook(w, r)
		}

	} else if len(SplitedPath) == 3 {

		if r.Method == "GET" {
			SearchBook(w, r)
		} else if r.Method == "DELETE" {
			RemoveBook(w, r)
		} else if r.Method == "PUT" {
			UpdateBook(w, r)
		}

	} else {

		w.WriteHeader(404)

	}
}

func RouteConfigure() {
	http.HandleFunc("/", MainRoute)
	http.HandleFunc("/books", RouteBook)
	http.HandleFunc("/books/", RouteBook)
}

func ServerConfigure() {
	RouteConfigure()
	fmt.Println("Server running in the port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	ServerConfigure()
}

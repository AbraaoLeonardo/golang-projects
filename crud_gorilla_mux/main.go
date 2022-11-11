package main

// Import the library
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(201)

	body, erro := ioutil.ReadAll(r.Body)
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
	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Books)

}

// Read 1 item
func SearchBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	SplitedPath := mux.Vars(r)
	id, _ := strconv.Atoi(SplitedPath["ID"])

	for _, book := range Books {
		if book.Id == id {
			json.NewEncoder(w).Encode(book)
		}
	}

	w.WriteHeader(404)

}

// Update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	SplitedPath := mux.Vars(r)
	id, erro := strconv.Atoi(SplitedPath["ID"])

	if erro != nil {
		w.WriteHeader(404)
		return
	}

	body, erroBody := ioutil.ReadAll(r.Body)

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
	w.Header().Set("content-type", "application/json")
	SplitedPath := mux.Vars(r)
	id, erro := strconv.Atoi(SplitedPath["ID"])

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

func RouteConfigure(route *mux.Router) {
	route.HandleFunc("/", MainRoute)
	route.HandleFunc("/books", ShowBooks).Methods("GET")
	route.HandleFunc("/books", AddNewBook).Methods("POST")
	route.HandleFunc("/books/{ID}", SearchBook).Methods("GET")
	route.HandleFunc("/books/{ID}", RemoveBook).Methods("DELETE")
	route.HandleFunc("/books/{ID}", UpdateBook).Methods("PUT")
}

func ServerConfigure() {
	route := mux.NewRouter().StrictSlash(true)
	RouteConfigure(route)
	fmt.Println("Server running in the port 8080")
	log.Fatal(http.ListenAndServe(":8080", route))
}

func main() {
	ServerConfigure()
}

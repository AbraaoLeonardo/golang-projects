package main

// Import the library
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Create a object book
type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

// Set the valuer of book
var Books []Book = []Book{
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

// GET books
func ShowBooks(w http.ResponseWriter, r *http.Request) {
	// We saying the header have valuer of json
	w.Header().Set("content-type", "application/json")
	// Encoding the json, to show in the screen
	encoder := json.NewEncoder(w)
	encoder.Encode(Books)
}
func AddNewBook(w http.ResponseWriter, r *http.Request) {

}

func RouteBook(w http.ResponseWriter, r *http.Request) {
	// If the method is not GET, we will return none
	if r.Method == "GET" {
		ShowBooks(w, r)
	} else if r.Method == "POST" {
		AddNewBook(w, r)
	}
}

func RouteConfigure() {
	http.HandleFunc("/", MainRoute)
	http.HandleFunc("/books", RouteBook)
}

func ServerConfigure() {
	RouteConfigure()
	fmt.Println("Server running in the port 8080")
	// If the port is used, we will recive a error
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	ServerConfigure()
}

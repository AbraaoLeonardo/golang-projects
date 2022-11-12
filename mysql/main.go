package main

// Import the library
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

// Create a object book
type BookModel struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Main Route function
func MainRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, 世界")
}

// Create book
func AddNewBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		fmt.Fprintf(w, "Error: %v", erro)
	}

	var newBook BookModel
	json.Unmarshal(body, &newBook)

	invalidData := VerifyData(newBook)

	if len(invalidData) > 0 {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(invalidData)
		return
	}
	_, errorCreate := db.Exec("INSERT INTO livros(author, title) Values(?, ?);", newBook.Author, newBook.Title)

	if errorCreate != nil {
		log.Println(errorCreate.Error())
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(newBook)
}

func VerifyData(book BookModel) string {
	if len(book.Author) == 0 || len(book.Author) >= 50 {
		return "Autor invalido"
	} else if len(book.Title) == 0 || len(book.Title) >= 50 {
		return "Titulo invalido"
	}
	return ""

}

// Read all item
func ShowBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	data, errorQuerry := db.Query("SELECT id, author, title FROM livros;")

	if errorQuerry != nil {
		fmt.Println(errorQuerry.Error())
		w.WriteHeader(500)
		return
	}

	var dbBooks []BookModel = make([]BookModel, 0)
	for data.Next() {
		var dataBook BookModel
		errorDataBook := data.Scan(&dataBook.Id, &dataBook.Author, &dataBook.Title)

		if errorDataBook != nil {
			fmt.Println("Book don't find " + errorDataBook.Error())
			continue
		}

		dbBooks = append(dbBooks, dataBook)
	}

	errorCloseDB := data.Close()
	if errorCloseDB != nil {
		fmt.Println("Error in close db " + errorCloseDB.Error())
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(dbBooks)

}

// Read 1 item
func SearchBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	SplitedPath := mux.Vars(r)
	id, _ := strconv.Atoi(SplitedPath["ID"])

	QueriedBook := db.QueryRow("SELECT * FROM livros WHERE id = ?", id)
	var dbBook BookModel
	errorBook := QueriedBook.Scan(&dbBook.Id, &dbBook.Author, &dbBook.Title)

	if errorBook != nil {
		fmt.Println("Error in search book: " + errorBook.Error())
		w.WriteHeader(400)
		return
	}
	json.NewEncoder(w).Encode(dbBook)

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

	QueriedBook := db.QueryRow("SELECT * FROM livros WHERE id = ?", id)
	var dbBook BookModel
	errorBook := QueriedBook.Scan(&dbBook.Id, &dbBook.Author, &dbBook.Title)

	if errorBook != nil {
		w.WriteHeader(400)
		fmt.Println("Error in query: " + errorBook.Error())
		return
	}

	_, errorUpdate := db.Exec("UPDATE livros SET author = ?, title =? WHERE id = ?;", modifiedBook.Author, modifiedBook.Title, id)

	if errorUpdate != nil {
		w.WriteHeader(400)
		fmt.Println("Error in update: " + errorUpdate.Error())
		return
	}

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

	QueriedBook := db.QueryRow("SELECT * FROM livros WHERE id = ?", id)
	var dbBook BookModel
	queryError := QueriedBook.Scan(&dbBook.Id, &dbBook.Author, &dbBook.Title)

	if queryError != nil {
		fmt.Println("Error in the query: " + queryError.Error())
		w.WriteHeader(500)
		return
	}

	_, errorDelete := db.Exec("DELETE FROM livros WHERE id = ?;", id)

	if errorDelete != nil {
		log.Println(errorDelete.Error())
		w.WriteHeader(400)
	}
}

func RouteConfigure(route *mux.Router) {
	route.HandleFunc("/", MainRoute)
	route.HandleFunc("/books/", ShowBooks).Methods("GET")
	route.HandleFunc("/books/", AddNewBook).Methods("POST")
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

func ConfigureDatabase() {
	var errorConnectDB error

	var user string = os.Getenv("DB_USER")
	var password string = os.Getenv("DB_PASSWORD")
	var route string = os.Getenv("DB_ROUTE")
	var database string = os.Getenv("DB_DATABASE")

	var dbCredential string = fmt.Sprintln("%s:%s@tcp(%s)/%s", user, password, route, database)

	db, errorConnectDB = sql.Open("mysql", dbCredential)
	if errorConnectDB != nil {
		log.Fatal(errorConnectDB.Error())
	}
}

func main() {
	ConfigureDatabase()
	ServerConfigure()
}

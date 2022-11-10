package main

// Import the libraries
import (
	"fmt"
	"log"
	"net/http"
)

// create the Handle of /form
func FormHandle(w http.ResponseWriter, r *http.Request) {
	// make the error output
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v ", err)
	}
	// Past the form information in the /form
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// Create the handle function of /hello
func HelloHandle(w http.ResponseWriter, r *http.Request) {
	// if the page != hello, we return error 404
	if r.URL.Path != "/hello" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	// if the method != GET, return method not supported
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	// What will be showed in the page
	fmt.Fprintf(w, "Hello")
}

func main() {
	// Create the server folder
	fileServer := http.FileServer(http.Dir("./static"))
	// Create the routes
	http.Handle("/", fileServer)
	http.HandleFunc("/form", FormHandle)
	http.HandleFunc("/hello", HelloHandle)

	//
	fmt.Print("Server running in the port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

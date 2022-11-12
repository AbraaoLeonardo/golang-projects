package main

import (
	"log"
	"net/http"

	"github.com/AbraaoLeonardo/golang-projects/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RouterBookStorage(r)
	r.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9000", r))
}

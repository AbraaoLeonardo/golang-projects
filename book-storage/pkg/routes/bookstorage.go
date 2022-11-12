package routes

import (
	"github.com/AbraaoLeonardo/golang-projects/pkg/controller"
	"github.com/gorilla/mux"
)

var RouterBookStorage = func(router *mux.Router) {
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controller.UpdateBook).Methods("GET")
	router.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
}

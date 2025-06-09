package main

import (
	"fmt"
	"log"
	"net/http"
	"restAPI/pkg/db"
	"restAPI/pkg/handlers"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the cardify REST API!")
	fmt.Println("cardify REST API")
}

func handleRequests() {
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", handlers.GetAllUsers).Methods(http.MethodGet)
	myRouter.HandleFunc("/users/{uid}", handlers.GetUser).Methods(http.MethodGet)
	myRouter.HandleFunc("/users", handlers.AddUser).Methods(http.MethodPost)
	myRouter.HandleFunc("/users/{uid}", handlers.UpdateUser).Methods(http.MethodPut)
	myRouter.HandleFunc("/users/{uid}", handlers.DeleteUser).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	DB := db.Connect()
	db.CreateTable(DB)
	handleRequests()
	db.CloseConnection(DB)
}

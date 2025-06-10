package main

import (
	"database/sql"
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

func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func handleRequests(DB *sql.DB) {
	h := handlers.New(DB)
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", h.GetAllUsers).Methods(http.MethodGet)
	myRouter.HandleFunc("/users/{uid}", h.GetUser).Methods(http.MethodGet)
	myRouter.HandleFunc("/users", h.AddUser).Methods(http.MethodPost)
	myRouter.HandleFunc("/users/{uid}", h.UpdateUser).Methods(http.MethodPut)
	myRouter.HandleFunc("/users/{uid}", h.DeleteUser).Methods(http.MethodDelete)

	handlerWithCors := enableCORS(myRouter)

	log.Fatal(http.ListenAndServe(":8080", handlerWithCors))
}

func main() {
	DB := db.Connect()
	db.CreateTable(DB)
	handleRequests(DB)
	db.CloseConnection(DB)
}

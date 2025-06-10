package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"restAPI/pkg/mocks"
	"restAPI/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	queryStmt := `SELECT * FROM users WHERE uid = $1 ;`
	results, err := h.DB.Query(queryStmt, uid)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	var user models.User
	for results.Next() {
		err = results.Scan(&user.Uid, &user.Username, &user.Email, &user.Picture)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(500)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	for _, user := range mocks.User {
		if user.Uid == uid {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
			break
		}
	}
}

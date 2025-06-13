package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"restAPI/pkg/mocks"
	"restAPI/pkg/models"

	"github.com/lib/pq"
)

func (h handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	results, err := h.DB.Query("SELECT * FROM users;")
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	var users = make([]models.User, 0)
	for results.Next() {
		var user models.User
		err = results.Scan(&user.Uid, &user.Username, &user.Email, &user.Picture, pq.Array(&user.Following), pq.Array(&user.Friends))
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(500)
			return
		}

		users = append(users, user)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.User)
}

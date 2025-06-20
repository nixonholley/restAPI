package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"restAPI/pkg/mocks"
	"restAPI/pkg/models"

	"github.com/gorilla/mux"
)

// type UserUpdate struct {
// 	Username string `json:"Username"`
// 	Email    string `json:"Email"`
// 	Picture  string `json:"Picture"`
// }

func (h handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	// Read request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedUser models.User
	json.Unmarshal(body, &updatedUser)

	queryStmt := `UPDATE users SET username = $2, email = $3, picture = $4 WHERE uid = $1 RETURNING uid;`
	err = h.DB.QueryRow(queryStmt, &uid, &updatedUser.Username, &updatedUser.Email, &updatedUser.Picture).Scan(&uid)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	// Read request body
	defer r.Body.Close()
	// log.Printf("Vars: %+v", vars)
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedUser models.User
	json.Unmarshal(body, &updatedUser)
	// if err != nil {
	// 	log.Println("Unmarshal error:", err)
	// }
	// log.Printf("Parsed body: %+v", updatedUser)

	for index, user := range mocks.User {
		if user.Uid == uid {
			user.Username = updatedUser.Username
			user.Email = updatedUser.Email
			user.Picture = updatedUser.Picture

			mocks.User[index] = user

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}

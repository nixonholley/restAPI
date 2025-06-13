package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"restAPI/pkg/mocks"
	"restAPI/pkg/models"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (h handler) AddUser(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(500)
		return
	}
	var user models.User
	json.Unmarshal(body, &user)

	queryStmt := `INSERT INTO users (uid,username,email,picture,following,friends) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uid;`
	err = h.DB.QueryRow(queryStmt, &user.Uid, &user.Username, &user.Email, &user.Picture, pq.Array(user.Following), pq.Array(user.Friends)).Scan(&user.Uid)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var user models.User
	json.Unmarshal(body, &user)

	user.Uid = (uuid.New()).String()
	mocks.User = append(mocks.User, user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

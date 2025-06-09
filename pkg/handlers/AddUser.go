package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"restAPI/pkg/mocks"
	"restAPI/pkg/models"

	"github.com/google/uuid"
)

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

package handlers

import (
	"encoding/json"
	"net/http"

	"restAPI/pkg/mocks"
)

func (h handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.User)
}

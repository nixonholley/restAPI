package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"restAPI/pkg/mocks"

	"github.com/gorilla/mux"
)

func (h handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	queryStmt := `DELETE FROM users WHERE uid = $1;`
	_, err := h.DB.Query(queryStmt, &uid)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	for index, user := range mocks.User {
		if user.Uid == uid {
			mocks.User = append(mocks.User[:index], mocks.User[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}

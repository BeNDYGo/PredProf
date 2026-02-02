package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"

	"predprof/databases/usersDatabase"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user usersDatabase.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if usersDatabase.UserExists(user.Username) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": "user already exists"})
		return
	}

	err = usersDatabase.CreateUser(user.Username, user.Password, 1000)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user created successfully"})
	fmt.Println("New user:", user.Username, user.Password)
}
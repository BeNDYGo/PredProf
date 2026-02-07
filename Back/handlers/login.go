package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"predprof/databases/usersDatabase"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user usersDatabase.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !usersDatabase.UserExists(user.Username) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	}

	valid, err := usersDatabase.CheckPassword(user.Username, user.Password)
	if err != nil || !valid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid password"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user logged in successfully"})
	fmt.Println("User logged in:", user.Username)
}

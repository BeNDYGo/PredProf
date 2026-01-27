package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users = make(map[string]User)

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		Users[user.Username] = user
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
		fmt.Println("User:", user.Username, "Password:", user.Password)
		return
	}

	w.Write([]byte("Register"))
}

func main() {
	http.HandleFunc("/register", register)
	fmt.Println("Server starting")
	http.ListenAndServe(":8080", nil)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"predprof/databases/usersDatabase"
)

func GetAllUserInfo(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	user, err := usersDatabase.GetUserAllInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

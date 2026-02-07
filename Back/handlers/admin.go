package handlers

import (
	"encoding/json"
	"net/http"
	"predprof/databases/usersDatabase"
)

func ChangeUserRole(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	role := r.URL.Query().Get("role")

	if username == "" || role == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "missing username or role"})
		return
	}

	if role != "admin" && role != "student" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid role"})
		return
	}

	if !usersDatabase.UserExists(username) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	err := usersDatabase.UpdateUserRole(username, role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to update role"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok", "role": role})
}

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

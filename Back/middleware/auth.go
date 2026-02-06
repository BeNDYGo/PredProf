package middleware

import (
    "net/http"
    "predprof/databases/usersDatabase"
	"encoding/json"
)

func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        username := r.URL.Query().Get("username")
        if username == "" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string {"error": "not username"})
            return
        }
        
        user, err := usersDatabase.GetUser(username)
        if err != nil || user.Role != "admin" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string {"error": "not admin"})
            return
        }
        
        next(w, r)
    }
}

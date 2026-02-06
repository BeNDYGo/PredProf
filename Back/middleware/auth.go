package middleware

import (
	"encoding/json"
	"net/http"
	"predprof/databases/usersDatabase"
)

func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем username из заголовка Authorization
		username := r.Header.Get("X-Username")
		if username == "" {
			// Пробуем получить из cookie
			cookie, err := r.Cookie("username")
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "not authenticated"})
				return
			}
			username = cookie.Value
		}

		user, err := usersDatabase.GetUser(username)
		if err != nil || user.Role != "admin" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{"error": "not admin"})
			return
		}

		next(w, r)
	}
}

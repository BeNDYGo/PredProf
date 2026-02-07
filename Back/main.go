package main

import (
	"fmt"
	"net/http"

	"predprof/databases/tasksDatabase"
	"predprof/databases/usersDatabase"
	"predprof/handlers"
	"predprof/middleware"
)

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Username")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {

	if err := usersDatabase.InitDB(); err != nil {
		panic(err)
	}
	if err := tasksDatabase.InitDB(); err != nil {
		panic(err)
	}

	http.HandleFunc("/api/register", corsMiddleware(handlers.Register))
	http.HandleFunc("/api/login", corsMiddleware(handlers.Login))
	http.HandleFunc("/api/getAllTasks", corsMiddleware(handlers.GetAllTasks))
	http.HandleFunc("/api/addTask", corsMiddleware(middleware.AdminOnly(handlers.AddTask)))
	http.HandleFunc("/api/userInfo", corsMiddleware(handlers.GetUserInfo))
	http.HandleFunc("/api/getUserAllInfo", corsMiddleware(middleware.AdminOnly(handlers.GetAllUserInfo)))
	http.HandleFunc("/api/changeRole", corsMiddleware(middleware.AdminOnly(handlers.ChangeUserRole)))
	http.HandleFunc("/api/ws", handlers.WsHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) })

	fmt.Println("Server starting")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

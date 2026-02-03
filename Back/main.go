package main

import (
	"fmt"
	"net/http"

	"predprof/databases/tasksDatabase"
	"predprof/databases/usersDatabase"
	"predprof/handlers"

)

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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
	

	http.HandleFunc("/register", corsMiddleware(handlers.Register))
	http.HandleFunc("/login", corsMiddleware(handlers.Login))
	http.HandleFunc("/getTasks", corsMiddleware(handlers.GetTasks))
	http.HandleFunc("/addTask", corsMiddleware(handlers.AddTask))
	http.HandleFunc("/userInfo", corsMiddleware(handlers.GetUserInfo))
	http.HandleFunc("/send", handlers.HandleWebSocket)

	fmt.Println("Server starting")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

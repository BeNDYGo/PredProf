package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"predprof/databases/usersDatabase"

	"github.com/gorilla/websocket"
)

type Message struct {
	user usersDatabase.User
	ans string
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "WebSocket upgrade failed", http.StatusBadRequest)
		return
	}

	defer conn.Close()

	for {
		var message Message

		if err := conn.ReadJSON(&message); err != nil {
			http.Error(w, "WebSocket read failed", http.StatusBadRequest)
			return
		}

		fmt.Println("WebSocker message", message.user, message.ans)
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	var username = r.URL.Query().Get("username")
	user, err := usersDatabase.GetUser(username)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"predprof/databases/tasksDatabase"
	"predprof/databases/usersDatabase"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn  *websocket.Conn
	match *Match
}

type Match struct {
	players map[*Client]bool
	answer  string
}

// сервер хранит только одного ожидающего игрока
var waiting *Client
var mtx sync.Mutex

func WsHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrader fallied", err)
		return
	}

	client := &Client{conn: conn}
	err = joinMatch(client)
	if err != nil {
		fmt.Println("joinMatch fallied", err)
		return
	}

	defer func() {
		leaveMatch(client)
		conn.Close()
	}()

	for {
		var msg map[string]interface{}
		if err := conn.ReadJSON(&msg); err != nil {
			fmt.Println("read JSON fallied", err)
			return
		}

		if client.match != nil {
			if client.match.answer == msg["userAnswer"] {
				// Уведомляем об итогах
				for player := range client.match.players {
					if player == client {
						player.conn.WriteJSON(map[string]string{"message": "you win"})
					} else {
						player.conn.WriteJSON(map[string]string{"message": "you lose"})
					}
				}
				return
			} else {
				client.conn.WriteJSON(map[string]string{"message": "incorrect"})
			}
		}
	}
}

func joinMatch(client *Client) error {
	mtx.Lock()
	defer mtx.Unlock()
	if waiting == nil {
		waiting = client
		client.conn.WriteJSON(map[string]string{"message": "waiting opponent..."})
		return nil
	}

	// Get task for the match
	task := tasksDatabase.GetTask("rus")

	match := Match{
		players: make(map[*Client]bool),
		answer:  task.Answer,
	}
	match.players[waiting] = true
	match.players[client] = true

	waiting.match = &match
	client.match = &match

	waiting.conn.WriteJSON(map[string]string{"message": "match found"})
	client.conn.WriteJSON(map[string]string{"message": "match found"})

	// Отправка задачи
	waiting.conn.WriteJSON(map[string]string{"task": task.Task})
	client.conn.WriteJSON(map[string]string{"task": task.Task})

	fmt.Println("match created")

	waiting = nil
	return nil
}

func leaveMatch(client *Client) {
	mtx.Lock()
	defer mtx.Unlock()
	if waiting == client {
		waiting = nil
	}
	if client.match != nil {
		for player := range client.match.players {
			if player != client {
				player.conn.WriteJSON(map[string]string{"message": "opponent disconected"})
				player.match = nil
			}
		}
		client.match = nil
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

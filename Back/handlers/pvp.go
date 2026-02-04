package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
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
}

var waiting *Client
var mtx sync.Mutex

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrader fallied", err)
		return
	}

	client := &Client{conn: conn}
	joinMatch(client)

	defer func() {
		leaveMatch(client)
		conn.Close()
	}()
	
	for {
		var msg string
		if err := conn.ReadJSON(&msg); err != nil {
			fmt.Println("read JSON fallied", err)
			return
		}
		for player := range client.match.players {
			if player != client {
				player.conn.WriteJSON(msg)
			}
		}
	}
}

func joinMatch(client *Client) {
	mtx.Lock()
	defer mtx.Unlock()
	if waiting == nil {
		waiting = client
		client.conn.WriteJSON("waiting opponent...")
		return
	}
	match := Match{
		players: make(map[*Client]bool),
	}
	match.players[waiting] = true
	match.players[client] = true

	waiting.match = &match
	client.match = &match

	waiting.conn.WriteJSON("match found")
	client.conn.WriteJSON("match found")

	fmt.Println("match created")

	waiting = nil

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
				player.conn.WriteJSON("opponent disconected")
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

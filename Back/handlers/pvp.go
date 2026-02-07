package handlers

import (
	"encoding/json"
	"fmt"
	"math"
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
	conn     *websocket.Conn
	match    *Match
	username string
}

type Match struct {
	players map[*Client]bool
	answer  string
}

// сервер хранит только одного ожидающего игрока
var waiting *Client
var mtx sync.Mutex

// calcElo рассчитывает новые рейтинги по формуле Elo (K=32)
func calcElo(winnerRating, loserRating int) (newWinner, newLoser int) {
	eWinner := 1.0 / (1.0 + math.Pow(10, float64(loserRating-winnerRating)/400.0))
	eLoser := 1.0 - eWinner
	K := 32.0
	newWinner = winnerRating + int(K*(1.0-eWinner))
	newLoser = loserRating + int(K*(0.0-eLoser))
	return
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrader fallied", err)
		return
	}

	client := &Client{conn: conn, username: username}
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
			fmt.Println("match closed")
			return
		}

		if client.match != nil {
			if client.match.answer == msg["userAnswer"] {
				// Определяем победителя и проигравшего
				winner := client
				var loser *Client
				for player := range client.match.players {
					if player != client {
						loser = player
					}
				}

				// Получаем текущие рейтинги и пересчитываем по Elo
				winnerUser, _ := usersDatabase.GetUser(winner.username)
				loserUser, _ := usersDatabase.GetUser(loser.username)
				newWinnerRating, newLoserRating := calcElo(winnerUser.Rating, loserUser.Rating)

				// Сохраняем в БД
				usersDatabase.UpdateAfterMatch(winner.username, newWinnerRating, true)
				usersDatabase.UpdateAfterMatch(loser.username, newLoserRating, false)

				// Уведомляем игроков с новыми рейтингами
				winner.conn.WriteJSON(map[string]interface{}{
					"message":   "you win",
					"newRating": newWinnerRating,
				})
				loser.conn.WriteJSON(map[string]interface{}{
					"message":   "you lose",
					"newRating": newLoserRating,
				})
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

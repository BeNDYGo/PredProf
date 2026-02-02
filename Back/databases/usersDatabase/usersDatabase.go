package usersDatabase

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var usersDB *sql.DB

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func InitDB() error {
	var err error

	usersDB, err = sql.Open("sqlite3", "./databases/usersDatabase/users.db")
	if err != nil {
		return err
	}

	_, err = usersDB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			username TEXT PRIMARY KEY,
			password TEXT,
			rating   INT
		)
	`)
	if err != nil {
		return err
	}
	return err
}

func UserExists(username string) bool {
	row := usersDB.QueryRow("SELECT username FROM users WHERE username = ?", username)
	var u string
	err := row.Scan(&u)
	return err == nil
}

func CreateUser(username string, password string, rating int64) error {
	_, err := usersDB.Exec(
		"INSERT INTO users(username, password) VALUES(?, ?)",
		username,
		password,
	)
	return err
}

func GetRating(username string) (int64, error) {
	row := usersDB.QueryRow("SELECT rating FROM users WHERE username = ?", username)
	var rating int64
	err := row.Scan(&rating)
	if err != nil {
		return 0, err
	}
	return rating, nil
}

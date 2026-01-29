package usersDatabase

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			username TEXT PRIMARY KEY,
			password TEXT
		)
	`)
	return err
}

func UserExists(username string) bool {
	row := db.QueryRow("SELECT username FROM users WHERE username = ?", username)
	var u string
	err := row.Scan(&u)
	return err == nil
}

func CreateUser(username, password string) error {
	_, err := db.Exec(
		"INSERT INTO users(username, password) VALUES(?, ?)",
		username,
		password,
	)
	return err
}

func GetDB() *sql.DB {
	return db
}

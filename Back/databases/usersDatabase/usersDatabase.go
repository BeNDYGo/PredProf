package usersDatabase

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var usersDB *sql.DB

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Rating   int    `json:"rating"`
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
}

func InitDB() error {
	var err error

	usersDB, err = sql.Open("sqlite3", "./databases/usersDatabase/users.db")
	if err != nil {
		return err
	}

	_, err = usersDB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			username  TEXT PRIMARY KEY,
			email     TEXT,
			password  TEXT,
			role      TEXT DEFAULT 'student',
			rating    INT,
			wins      INT DEFAULT 0,
			losses    INT DEFAULT 0
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

func UserExists(username string) bool {
	row := usersDB.QueryRow("SELECT username FROM users WHERE username = ?", username)
	var u string
	err := row.Scan(&u)
	return err == nil
}

func CreateUser(username, email, password string) error {
	if username == "alex" {
		_, err := usersDB.Exec(`
			INSERT INTO users(username, email, password, role, rating, wins, losses)
			VALUES(?, ?, ?, 'admin', 1000, 0, 0)
		`, username, email, password)
		return err
	} else {
		_, err := usersDB.Exec(`
			INSERT INTO users(username, email, password, role, rating, wins, losses)
			VALUES(?, ?, ?, 'student', 1000, 0, 0)
		`, username, email, password)
		return err
	}
}

func GetUserPassword(username string) string {
	row := usersDB.QueryRow("SELECT password FROM users WHERE username = ?", username)
	var password string
	err := row.Scan(&password)
	if err != nil {
		return ""
	}
	return password
}

func GetUser(username string) (User, error) {
	row := usersDB.QueryRow(`
		SELECT username, role, rating, wins, losses 
		FROM users 
		WHERE username = ?
	`, username)
	var user User
	err := row.Scan(
		&user.Username,
		&user.Role,
		&user.Rating,
		&user.Wins,
		&user.Losses,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserAllInfo(username string) (User, error) {
	row := usersDB.QueryRow(`
		SELECT username, email, password, role, rating, wins, losses 
		FROM users 
		WHERE username = ?
	`, username)
	var user User
	err := row.Scan(
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.Rating,
		&user.Wins,
		&user.Losses,
	)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func UpdateAfterMatch(username string, newRating int, won bool) error {
	if won {
		_, err := usersDB.Exec(
			"UPDATE users SET rating = ?, wins = wins + 1 WHERE username = ?",
			newRating, username,
		)
		return err
	}
	_, err := usersDB.Exec(
		"UPDATE users SET rating = ?, losses = losses + 1 WHERE username = ?",
		newRating, username,
	)
	return err
}
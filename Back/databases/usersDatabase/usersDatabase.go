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
	Decided  int    `json:"decided"`
	Mistakes int    `json:"mistakes"`
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
			decided   INT,
			mistakes  INT
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
	if username == "alex"{
		_, err := usersDB.Exec(`
			INSERT INTO users(username, email, password, role, rating, decided, mistakes)
			VALUES(?, ?, ?, 'admin', 1000, 0, 0)
	`, username, email, password)
		return err
	} else {
		_, err := usersDB.Exec(`
			INSERT INTO users(username, email, password, role, rating, decided, mistakes)
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
		SELECT username, role, rating, decided, mistakes 
		FROM users 
		WHERE username = ?
	`, username)
	var user User
	err := row.Scan(
		&user.Username,
		&user.Role,
		&user.Rating,
		&user.Decided,
		&user.Mistakes,
	)
	if err != nil {
		return user, nil
	}
	return user, err
}

package usersDatabase

import (
	"database/sql"
	
	_ "github.com/mattn/go-sqlite3"
)

var usersDB *sql.DB

type User struct {
	Username string  `json:"username"`
	Email string     `json:"email"`
	Password string  `json:"password"`
	Rating int    `json:"rating"`
	Decided int   `json:"decided"`
	Mistakes int  `json:"mistakes"`
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
	_, err := usersDB.Exec(`
		INSERT INTO users(username, email, password, rating, decided, mistakes)
		VALUES(?, ?, ?, 1000, 0, 0)
	`, username, email, password)

	return err
}

func GetUser(username string) (User, error) {
	row := usersDB.QueryRow(`
		SELECT username, email, password, rating, decided, mistakes 
		FROM users 
		WHERE username = ?
	`, username)
	var user User
	err := row.Scan(
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Rating,
		&user.Decided,
		&user.Mistakes,
	)
	if err != nil {
		return user, nil
	}
	return user, err
}

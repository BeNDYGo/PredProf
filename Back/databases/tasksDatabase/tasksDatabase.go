package tasksDatabase

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Task       string `json:"task"`
	Answer     string `json:"answer"`
	TaskType   string `json:"taskType,omitempty"`
	Difficulty string `json:"difficulty,omitempty"`
}

var tasksRusDB *sql.DB
var tasksMathDB *sql.DB

func InitDB() error {
	var err error

	tasksRusDB, err = sql.Open("sqlite3", "./databases/tasksDatabase/tasks_rus.db")
	if err != nil {
		return err
	}

	_, err = tasksRusDB.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			task TEXT NOT NULL,
			answer TEXT NOT NULL,
			taskType TEXT,
			difficulty TEXT
		)
	`)
	if err != nil {
		return err
	}

	tasksMathDB, err = sql.Open("sqlite3", "./databases/tasksDatabase/tasks_math.db")
	if err != nil {
		return err
	}

	_, err = tasksMathDB.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			task TEXT NOT NULL,
			answer TEXT NOT NULL,
			taskType TEXT,
			difficulty TEXT
		)
	`)
	return err
}

func GetAllTasks(subject string, taskType string, difficulty string) ([]Task, error) {
	var targetDB *sql.DB
	switch subject {
	case "rus":
		targetDB = tasksRusDB
	case "math":
		targetDB = tasksMathDB
	default:
		return nil, fmt.Errorf("unknown subject: %s", subject)
	}

	query := "SELECT task, answer, taskType, difficulty FROM tasks WHERE 1=1"
	args := []interface{}{}

	if taskType != "" && taskType != "none" {
		query += " AND taskType = ?"
		args = append(args, taskType)
	}

	if difficulty != "" && difficulty != "none" {
		query += " AND difficulty = ?"
		args = append(args, difficulty)
	}

	rows, err := targetDB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Task, &task.Answer, &task.TaskType, &task.Difficulty)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTask(subject string){}

func AddTask(subject string, task string, answer string, taskType string, difficulty string) error {
	var targetDB *sql.DB
	switch subject {
	case "rus":
		targetDB = tasksRusDB
	case "math":
		targetDB = tasksMathDB
	default:
		return fmt.Errorf("unknown subject: %s", subject)
	}

	_, err := targetDB.Exec(
		"INSERT INTO tasks(task, answer, taskType, difficulty) VALUES(?, ?, ?, ?)",
		task,
		answer,
		taskType,
		difficulty,
	)
	return err
}

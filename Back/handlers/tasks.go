package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"

	"predprof/databases/tasksDatabase"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	subject := r.URL.Query().Get("subject")
	if subject == "" {
		http.Error(w, "subject parameter is required", http.StatusBadRequest)
		return
	}

	taskType := r.URL.Query().Get("taskType")
	difficulty := r.URL.Query().Get("difficulty")

	tasks, err := tasksDatabase.GetAllTasks(subject, taskType, difficulty)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var task tasksDatabase.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subject := r.URL.Query().Get("subject")
	if subject == "" {
		http.Error(w, "subject parameter is required", http.StatusBadRequest)
		return
	}

	err = tasksDatabase.AddTask(subject, task.Task, task.Answer, task.TaskType, task.Difficulty)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "task added"})
	fmt.Println("Task added")
}
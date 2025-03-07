package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"taskapi/internal/model"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	includeDeleted := r.URL.Query().Get("deleted") == "true"

	query := `SELECT id, task, created_at, deleted_at, deleted FROM tasks`
	if !includeDeleted {
		query += ` WHERE deleted = false`
	}

	rows, err := h.db.Query(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		err := rows.Scan(&t.ID, &t.Task, &t.CreatedAt, &t.DeletedAt, &t.Deleted)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if task.Task == "" {
		http.Error(w, "Task cannot be empty", http.StatusBadRequest)
		return
	}

	var newTask model.Task
	err := h.db.QueryRow(
		`INSERT INTO tasks (task) VALUES ($1) RETURNING id, task, created_at, deleted_at, deleted`,
		task.Task,
	).Scan(&newTask.ID, &newTask.Task, &newTask.CreatedAt, &newTask.DeletedAt, &newTask.Deleted)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	now := time.Now()
	result, err := h.db.Exec(
		`UPDATE tasks SET deleted = true, deleted_at = $1 WHERE id = $2 AND deleted = false`,
		now, id,
	)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Task not found or already deleted", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

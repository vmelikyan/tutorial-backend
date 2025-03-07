package api

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(h *Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", h.HealthCheck).Methods("GET")
	r.HandleFunc("/tasks", h.GetTasks).Methods("GET")
	r.HandleFunc("/task", h.CreateTask).Methods("POST")
	r.HandleFunc("/task/{id}", h.DeleteTask).Methods("DELETE")

	return r
}

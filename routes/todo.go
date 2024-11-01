package routes

import (
	"encoding/json"
	"net/http"
	"todo/controllers"

	"github.com/go-chi/chi/v5"
)

func CheckServerWorking(w http.ResponseWriter, r *http.Request) {
	var message = map[string]string{"message": "Hello World Todo"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func TodoRoutes(r *chi.Mux) {
	r.Post("/v1", controllers.CreateTask)
	r.Get("/v1", controllers.GetAllTasks)
}

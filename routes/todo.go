package routes

import (
	"todo/controllers"

	"github.com/go-chi/chi/v5"
)

func TodoRoutes(r *chi.Mux) {
	r.Post("/v1", controllers.CreateTask)
	r.Get("/v1", controllers.GetAllTasks)
	r.Put("/v1", controllers.UpdateTask)
	r.Delete("/v1", controllers.DeleteTask)

}

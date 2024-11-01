package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"todo/db"
	"todo/models"
	"todo/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func checkServerWorking(w http.ResponseWriter, r *http.Request) {
	var message = map[string]string{"message": "Hello World!!!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	db.ConnectMongo()
	defer db.DisconnectMongo()
	client := db.GetDB()
	models.SetCollection(client.Database("go-todo").Collection("task"))
	routes.TodoRoutes(r)

	// r.Get("/", routes.CheckServerWorking)
	PORT := os.Getenv("PORT")
	fmt.Println("server started at ", PORT)
	if err := http.ListenAndServe(PORT, r); err != nil {
		fmt.Println("Error starting server", err)
		panic(err)

	}

}

package controllers

import (
	"encoding/json"
	"net/http"
	"todo/models"
)

type ErrorMessage struct {
	message string
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var todo models.TODO
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := models.CreateTODO(todo)
	if err != nil {
		var message = map[string]string{"message": "Something went wrong!!!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
		return
	}
	var successmessage = map[string]string{"message": "Task Created!!!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(successmessage)

}
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	result, err := models.GetTODO()
	if err != nil {
		var message = map[string]string{"message": "Something went wrong!!!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)

}

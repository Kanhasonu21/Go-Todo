package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"todo/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var todo models.TODO
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.CreatedAt = time.Now()
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

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var toUpdateTask models.TODO
	if err := json.NewDecoder(r.Body).Decode(&toUpdateTask); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	isSuccess, _ := models.UpdateTask(id, toUpdateTask)
	if isSuccess {
		var message = map[string]string{"message": "Task Updated!!!"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(message)
		return
	} else {
		var message = map[string]string{"message": "Something went wrong!!!"}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
		return
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	isSuccess, _ := models.DeleteTask(id)
	if isSuccess {
		var message = map[string]string{"message": "Task Deleted!!!"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(message)
		return
	} else {
		var message = map[string]string{"message": "Something went wrong!!!"}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
		return
	}
}

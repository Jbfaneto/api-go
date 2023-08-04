package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Jbfaneto/api-go/models"
)

// Create is a function to create a new todo it receives a http.ResponseWriter and a http.Request and returns nothing
func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Massage": fmt.Sprintf("Error: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserted with success! ID: %d", id),
		}
	}
	// set the header to application/json
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

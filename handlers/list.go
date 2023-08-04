package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jbfaneto/api-go/models"
)

// List is a function to list all todos it receives a http.ResponseWriter and a http.Request and returns nothing
func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

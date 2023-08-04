package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Jbfaneto/api-go/models"
	"github.com/go-chi/chi/v5"
)

// Get is a function to get a todo it receives a http.ResponseWriter and a http.Request and returns nothing
func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error ao fazer parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Error ao procurar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// set the header to application/json
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

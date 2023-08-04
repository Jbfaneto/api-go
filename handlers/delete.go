package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Jbfaneto/api-go/models"
	"github.com/go-chi/chi/v5"
)

// Delete is a function to delete a todo it receives a http.ResponseWriter and a http.Request and returns nothing
func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error ao fazer parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Error ao Deletar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: Foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Todo removido com sucesso! ID: %d", id),
	}
	// set the header to application/json
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

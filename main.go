package main

import (
	"fmt"
	"net/http"

	"github.com/Jbfaneto/api-go/configs"
	"github.com/Jbfaneto/api-go/handlers"
	"github.com/go-chi/chi/v5"
)

// the main function of the app
func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	// create a new router
	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.Get)
	r.Get("/", handlers.List)
	// start the server
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}

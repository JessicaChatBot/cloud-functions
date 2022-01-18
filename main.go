package main

import (
	"net/http"

	"github.com/JessicaChatBot/cloud-functions/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/say", api.GetHttpRequestProcessingFunction())
	http.ListenAndServe(":3000", r)
}

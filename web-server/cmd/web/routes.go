package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tatyanayagolnikov/go-website/pkg/config"
	"github.com/tatyanayagolnikov/go-website/pkg/handlers"
	//	"github.com/tatyanayagolnikov/go-website/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// middleware
	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/person", handlers.Repo.Person)
	return mux // mux - multiplexer or http handler
}


package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/leonardodk/HotelDK/pkg/config"
	"github.com/leonardodk/HotelDK/pkg/handlers"
)

// routes returns a mux with the handlers registered, ready to be used as an http.Handler
func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(LoadSession)

	mux.Get("/", handlers.PackageRepo.Home)
	mux.Get("/about", handlers.PackageRepo.About)

	return mux
}

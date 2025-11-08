package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func MyMiddleware(next http.Handler) http.Handler {

	// make a func doing some stuffg
	myFunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit the page")
		next.ServeHTTP(w, r)
	}
	// turn it into a handler and return it
	return http.HandlerFunc(myFunc)
}

// adds CSRF protection to all POST  requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// load session loads and saves the session on every request
func LoadSession(next http.Handler) http.Handler {
	return MainSession.LoadAndSave(next)
}

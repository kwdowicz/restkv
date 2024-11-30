package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func v1Router(a application) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(SetContentType("application/json"))
	r.Get("/", a.health)
	r.Get("/get", a.get)
	r.Post("/set", a.set)
	return r
}

func SetContentType(ct string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", ct)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

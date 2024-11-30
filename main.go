package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	store := NewKVStore()
	app := &application{store: store}
	r := chi.NewRouter()
	r.Mount("/v1", v1Router(*app))
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", r)
}

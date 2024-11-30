package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type storage interface {
	Set(string, string)
	Get(string) (string, bool)
}

type application struct {
	store storage
}

func (a *application) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v\n", a.store)))
	//w.Write([]byte("ok"))
}

func (a *application) set(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid request body", http.StatusMethodNotAllowed)
		return

	}
	a.store.Set(data.Key, data.Value)
	w.WriteHeader(http.StatusOK)
}

func (a *application) get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "key is required", http.StatusBadRequest)
	}
	value, _ := a.store.Get(key)
	json.NewEncoder(w).Encode(map[string]string{"value": value})
}

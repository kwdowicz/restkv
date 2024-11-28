package main

import (
	"encoding/json"
	"net/http"
)

func set(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var data struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid request body", http.StatusMethodNotAllowed)
		return

	}
	store.Set(data.Key, data.Value)
	w.WriteHeader(http.StatusOK)
}

func get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "key is required", http.StatusBadRequest)
	}
	value, exists := store.Get(key)
	if !exists {
		http.Error(w, "key not found", http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(map[string]string{"value": value})
}

package main

import (
	"fmt"
	"net/http"
)

var store *KVStore

func main() {
	store = NewKVStore()
	http.HandleFunc("/set", set) 
	http.HandleFunc("/get", get) 

	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"sync"
)

type KVStore struct {
	store map[string]string
	mux   sync.RWMutex
}

func NewKVStore() *KVStore {
	return &KVStore{
		store: make(map[string]string),
	}
}

func (kv *KVStore) Set(key, value string) {
	kv.mux.Lock()
	defer kv.mux.Unlock()
	kv.store[key] = value
}

func (kv *KVStore) Get(key string) (string, bool) {
	kv.mux.RLock()
	kv.mux.RUnlock()
	value, exists := kv.store[key]
	return value, exists
}

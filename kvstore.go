package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type KVStore struct {
	store   map[string]string
	mux     sync.RWMutex
	logFile string
}

func NewKVStore() *KVStore {
	kvstore := &KVStore{
		store:   make(map[string]string),
		logFile: "log.txt",
	}
	kvstore.load()
	return kvstore
}

func (kv *KVStore) Set(key, value string) {
	kv.mux.Lock()
	defer kv.mux.Unlock()
	if err := kv.log(key, value); err != nil {
		return
	}
	kv.store[key] = value
}

func (kv *KVStore) Get(key string) (string, bool) {
	kv.mux.RLock()
	kv.mux.RUnlock()
	value, exists := kv.store[key]
	return value, exists
}

func (kv *KVStore) GetMap() map[string]string {
	kv.mux.RLock()
	kv.mux.RUnlock()
	return kv.store 
}

func (kv *KVStore) log(key, value string) error {
	file, err := os.OpenFile(kv.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
		return err
	}
	defer file.Close()
	data := fmt.Sprintf("%s:%s\n", key, value)
	if _, err := file.WriteString(data); err != nil {
		log.Fatalf("failed to write to file: %v\n", err)
		return err
	}
	return nil
}

func (kv *KVStore) load() {
	if !fileExists(kv.logFile) {
		return
	}
	file, err := os.Open(kv.logFile)
	if err != nil {
		log.Fatalf("log file: %s exists but can't be open, error: %v", kv.logFile, err)
		panic("terminating...")
	}
	defer file.Close()

	kv.mux.Lock()
	defer kv.mux.Unlock()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), ":")
		key := splitted[0]
		value := ""
		if len(splitted) > 1 {
			value = splitted[1]
		}
		kv.store[key] = value
	}
}

func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

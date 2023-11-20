package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type connectionCallback func()

func handler(w http.ResponseWriter, r *http.Request, onConnect connectionCallback) {
	onConnect() // Invoke the callback function to log the connection event

	message := map[string]string{"message": "Hello, World!"}
	jsonData, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	connectionLogger := func() {
		log.Println("Client connected to port 8000")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, connectionLogger)
	})

	log.Println("Server listening on port 8000")
	http.ListenAndServe(":8000", nil)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// Prepare the data to be sent in the request body
	data := LoginData{
		Username: "admin",
		Password: "securepassword",
	}

	// Convert data to JSON
	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// Send a POST request to the server
	resp, err := http.Post("https://localhost:8080/login", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalf("Error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	// Print the server response
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	// Print the response message
	fmt.Printf("Response: %v\n", response)
}

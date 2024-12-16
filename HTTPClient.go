package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Make an HTTP GET request to the server
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body using io.ReadAll
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Print the response body
	fmt.Println("Response from server:", string(body))
}

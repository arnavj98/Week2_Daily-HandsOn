package main

import (
	"fmt"
	"net/http"
)

// Handler function for the root URL
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Writing a response to the client
	fmt.Fprintf(w, "Hello, Go HTTP Server!")
}

func main() {
	// Register the handler with the default HTTP mux
	http.HandleFunc("/", helloHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

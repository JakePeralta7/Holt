package main

import "net/http"

func main() {
	// Create a new ServeMux (router)
	mux := http.NewServeMux()

	// Define a route for the help ("/help") path
	mux.Handle("/help", http.HandlerFunc(handler_help))

	// Start the server on port 8080
	http.ListenAndServe(":8080", mux)
}

package main

import (
	"fmt"
	"net/http"
	"os"
)

var GLOBAL_SERVER_HOSTNAME string
var GLOBAL_CONFIG = make(map[string]map[string]interface{})

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
		os.Exit(1)
	}
	GLOBAL_SERVER_HOSTNAME = hostname

	// Loading configuration from disk (implemented in config.go)
	load_config()

	write_log("init", "Starting the server")

	// Create a new ServeMux (router)
	mux := http.NewServeMux()

	// Define a route for the help ("/help") path
	mux.Handle("/help", http.HandlerFunc(handler_help))

	// Start the server on port 8080
	http.ListenAndServe(":8080", mux)
}

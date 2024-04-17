package main

import (
	"fmt"
	"os"
)

var GLOBAL_CLIENT_HOSTNAME string
var GLOBAL_CONFIG = make(map[string]map[string]interface{})

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
		os.Exit(1)
	}
	GLOBAL_CLIENT_HOSTNAME = hostname

	// Loading configuration from disk (implemented in config.go)
	load_config()

	write_log("init", "Starting the client")

	get("/help")
}

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
		write_log("init", fmt.Sprintf("Error getting hostname: %v\n", err))
		os.Exit(1)
	}
	GLOBAL_CLIENT_HOSTNAME = hostname

	write_log("init", "Starting the client")

	// Loading configuration from disk (implemented in config.go)
	load_config()

	get("/help")
}

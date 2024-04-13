package main

import (
	"fmt"
	"os"
)

var global_holt_host string
var global_holt_port int
var global_holt_interval int

func main() {

	// Loading configuration from disk (implemented in config.go)
	config, err := load_config()
	if err != nil {
		write_log("init", fmt.Sprintf("Error loading config: %v\n", err))
		os.Exit(1)
	}

	// Printing the server section
	current_config_section := "server"
	if config[current_config_section] != nil {
		global_holt_host = config[current_config_section]["host"].(string)
		global_holt_port = config[current_config_section]["port"].(int)
		global_holt_interval = config[current_config_section]["interval"].(int)

		// Logs the loaded `server` configuration
		write_log("init", fmt.Sprintf("Loading %s configuration:\n\t[%s]\n\thost = %s\n\tport = %d\n\tinterval = %d\n",
			current_config_section, current_config_section, global_holt_host, global_holt_port, global_holt_interval))
	}

	get("/help")
}

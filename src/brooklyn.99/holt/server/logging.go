package main

import (
	"fmt"
	"log"
	"os"
)

func write_log(component string, message string) {

	log_path := fmt.Sprintf("%s\\server.log", GLOBAL_CONFIG["server"]["log_dir"])

	// Open a file for writing logs
	file, err := os.OpenFile(log_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Formatting the prefix for the line, for example "WIN10-HOST [init]:"
	prefix := fmt.Sprintf("%s [%s]: ", GLOBAL_SERVER_HOSTNAME, component)

	// Create a new logger
	logger := log.New(file, prefix, log.LstdFlags)

	// Write the log
	logger.Println(message)
}

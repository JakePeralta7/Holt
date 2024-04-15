package main

import (
	"fmt"
	"log"
	"os"
)

func write_log(component string, message string) {

	// Open a file for writing logs
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Formatting the prefix for the line, for example "WIN10-HOST [init]:"
	prefix := fmt.Sprintf("%s [%s]: ", GLOBAL_CLIENT_HOSTNAME, component)

	// Create a new logger
	logger := log.New(file, prefix, log.LstdFlags)

	// Write the log
	logger.Println(message)
}

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

	prefix := fmt.Sprintf("%s [%s]: ", global_client_hostname, component)

	// Create a new logger
	logger := log.New(file, prefix, log.LstdFlags)

	// Write the log
	logger.Println(message)
}

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

	component = fmt.Sprintf("%s: ", component)

	// Create a new logger
	logger := log.New(file, component, log.LstdFlags)

	// Write the log
	logger.Println(message)
}

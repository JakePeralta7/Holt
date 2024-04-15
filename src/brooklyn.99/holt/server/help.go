package main

import (
	"fmt"
	"net/http"
)

func handler_help(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hi there!\n")
	fmt.Fprintf(w, "Hello World")
}

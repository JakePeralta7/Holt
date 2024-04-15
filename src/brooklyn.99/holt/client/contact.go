package main

import (
	"fmt"
	"io"
	"net/http"
)

func get(uri string) {
	url := fmt.Sprintf("http://%s:%d%s",
		global_holt_host, global_holt_port, uri)

	fmt.Printf("URL: %s\n",
		url)

	resp, err := http.Get(url)

	if err != nil {
		write_log("contact", err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		write_log("contact", err.Error())
	}

	write_log("contact", fmt.Sprintf("Recieved '%s' from the server", string(body)))
}

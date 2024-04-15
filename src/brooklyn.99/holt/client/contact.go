package main

import (
	"fmt"
	"io"
	"net/http"
)

func get(uri string) {
	url := fmt.Sprintf("http://%s:%d%s", global_server_host, global_server_port, uri)

	write_log("contact", fmt.Sprintf("Performing GET request to '%s'", url))

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

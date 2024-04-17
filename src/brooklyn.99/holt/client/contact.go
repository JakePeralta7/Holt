package main

import (
	"fmt"
	"io"
	"net/http"
)

func get(uri string) {
	url := fmt.Sprintf("http://%s:%d%s", GLOBAL_CONFIG["server"]["host"], GLOBAL_CONFIG["server"]["port"], uri)

	write_log("contact", fmt.Sprintf("Performing GET request to '%s'", url))

	resp, err := http.Get(url)
	if err != nil {
		write_log("contact", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		write_log("contact", err.Error())
		return
	}

	write_log("contact", fmt.Sprintf("Recieved '%s' from the server", string(body)))
}

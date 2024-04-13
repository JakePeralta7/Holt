package main

import (
	"strconv"

	"gopkg.in/ini.v1"
)

func load_config() (map[string]map[string]interface{}, error) {

	// Loading the config from the path `$CWD\config.ini`
	cfg, err := ini.Load("config.ini")
	if err != nil {
		return nil, err
	}

	config := make(map[string]map[string]interface{})

	// Loading the `server` section from the config file
	server_section_name := "server"
	server_section, err := cfg.GetSection(server_section_name)
	if err == nil {
		config_server := make(map[string]interface{})

		// DNS or IP address of the Holt server
		config_server["host"] = server_section.Key("host").String()

		// The port the Holt server listens on
		server_port, _ := strconv.Atoi(server_section.Key("port").String())
		config_server["port"] = server_port

		// The interval in which the client will contact Holt
		server_interval, _ := strconv.Atoi(server_section.Key("interval").String())
		config_server["interval"] = server_interval

		config[server_section_name] = config_server
	}

	return config, nil
}

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

		// The instance name of the server
		server_instance_name := server_section.Key("instance_name").String()
		if server_instance_name != "" {
			config_server["instance_name"] = server_instance_name
		} else {
			config_server["instance_name"] = "dev-holt.example.com"
		}

		// DNS or IP address of the Holt server
		server_host := server_section.Key("host").String()
		if server_host != "" {
			config_server["host"] = server_host
		} else {
			config_server["host"] = "127.0.0.1"
		}

		// The port the Holt server listens on
		server_port, _ := strconv.Atoi(server_section.Key("port").String())
		if server_port != 0 {
			config_server["port"] = server_port
		} else {
			config_server["port"] = 8080
		}

		// The interval in which the client will contact Holt
		server_interval, _ := strconv.Atoi(server_section.Key("interval").String())
		if server_interval != 0 {
			config_server["interval"] = server_interval
		} else {
			config_server["interval"] = 15
		}

		config[server_section_name] = config_server
	}

	return config, nil
}

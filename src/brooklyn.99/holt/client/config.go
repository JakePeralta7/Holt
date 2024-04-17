package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func load_config() {

	// Loading the config from the path `$CWD\config.ini`
	cfg, err := ini.Load("config.ini")
	if err != nil {
		write_log("config", fmt.Sprintf("Error loading config: %v\n", err))
		os.Exit(1)
	}

	var current_section_name string

	// Loading the `client` section from the config file
	current_section_name = "client"
	client_section, err := cfg.GetSection(current_section_name)
	if err != nil {
		write_log("config", err.Error())
		os.Exit(1)
	} else {
		config_client := make(map[string]interface{})

		// The port the Holt client listens on
		client_port, err := client_section.Key("port").Int()
		if err != nil {
			config_client["port"] = 8080
		} else {
			config_client["port"] = client_port
		}

		// The directory in which the client.log will be stored
		client_log_dir := client_section.Key("log_dir").String()
		if client_log_dir != "" {
			config_client["log_dir"] = client_log_dir
		} else {
			config_client["log_dir"] = "."
		}

		GLOBAL_CONFIG[current_section_name] = config_client
	}

	// Logs the loaded `client` configuration
	write_log("config", fmt.Sprintf("Loading %s configuration:\n\t[%s]\n\tport = %d\n\tlog_dir = %s\n",
		current_section_name, current_section_name, GLOBAL_CONFIG[current_section_name]["port"], GLOBAL_CONFIG[current_section_name]["log_dir"]))

	// Loading the `server` section from the config file
	current_section_name = "server"
	server_section, err := cfg.GetSection(current_section_name)
	if err != nil {
		write_log("config", err.Error())
		os.Exit(1)
	} else {
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
		server_port, err := server_section.Key("port").Int()
		if err != nil {
			config_server["port"] = 8080
		} else {
			config_server["port"] = server_port
		}

		// The interval in which the client will contact Holt
		server_interval, err := server_section.Key("interval").Int()
		if err != nil {
			config_server["interval"] = 15
		} else {
			config_server["interval"] = server_interval
		}

		GLOBAL_CONFIG[current_section_name] = config_server
	}

	// Logs the loaded `server` configuration
	write_log("config", fmt.Sprintf("Loading %s configuration:\n\t[%s]\n\tinstance_name = %s\n\thost = %s\n\tport = %d\n\tinterval = %d\n",
		current_section_name, current_section_name, GLOBAL_CONFIG[current_section_name]["instance_name"], GLOBAL_CONFIG[current_section_name]["host"], GLOBAL_CONFIG[current_section_name]["port"], GLOBAL_CONFIG[current_section_name]["interval"]))
}

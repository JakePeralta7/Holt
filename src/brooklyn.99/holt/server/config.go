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

		// The directory in which the client.log will be stored
		client_log_dir := server_section.Key("log_dir").String()
		if client_log_dir != "" {
			config_server["log_dir"] = client_log_dir
		} else {
			config_server["log_dir"] = "."
		}

		GLOBAL_CONFIG[current_section_name] = config_server
	}

	// Logs the loaded `server` configuration
	write_log("config", fmt.Sprintf("Loading %s configuration:\n\t[%s]\n\tinstance_name = %s\n\thost = %s\n\tport = %d\n\tlog_dir = %s\n",
		current_section_name, current_section_name, GLOBAL_CONFIG[current_section_name]["instance_name"], GLOBAL_CONFIG[current_section_name]["host"], GLOBAL_CONFIG[current_section_name]["port"], GLOBAL_CONFIG[current_section_name]["log_dir"]))

	// Loading the `database` section from the config file
	current_section_name = "database"
	database_section, err := cfg.GetSection(current_section_name)
	if err != nil {
		write_log("config", err.Error())
		os.Exit(1)
	} else {
		config_database := make(map[string]interface{})

		// The instance name of the server
		server_instance_name := database_section.Key("instance_name").String()
		if server_instance_name != "" {
			config_database["instance_name"] = server_instance_name
		} else {
			config_database["instance_name"] = "dev-holt.example.com"
		}

		// DNS or IP address of the Holt server
		server_host := database_section.Key("host").String()
		if server_host != "" {
			config_database["host"] = server_host
		} else {
			config_database["host"] = "127.0.0.1"
		}

		// The port the Holt server listens on
		server_port, err := database_section.Key("port").Int()
		if err != nil {
			config_database["port"] = 8080
		} else {
			config_database["port"] = server_port
		}

		// The directory in which the client.log will be stored
		client_log_dir := database_section.Key("log_dir").String()
		if client_log_dir != "" {
			config_database["log_dir"] = client_log_dir
		} else {
			config_database["log_dir"] = "."
		}

		GLOBAL_CONFIG[current_section_name] = config_database
	}

	// Logs the loaded `server` configuration
	write_log("config", fmt.Sprintf("Loading %s configuration:\n\t[%s]\n\tinstance_name = %s\n\thost = %s\n\tport = %d\n\tlog_dir = %s\n",
		current_section_name, current_section_name, GLOBAL_CONFIG[current_section_name]["instance_name"], GLOBAL_CONFIG[current_section_name]["host"], GLOBAL_CONFIG[current_section_name]["port"], GLOBAL_CONFIG[current_section_name]["log_dir"]))
}

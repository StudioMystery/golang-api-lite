package utility

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Env                     string `json:"env"`
	Connection              string `json:"connection"`
	Foobar_FrontendHostname string `json:"foobar-frontend-hostname"`
	Foobar_BackendHostname  string `json:"foobar-backend-hostname"`
}

// Get the config struct so I can use the config.json values in code
func GetConfig() Config {
	// Read the configuration file
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse the JSON configuration
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	// Return Config
	return config
}

// Used to get full url and port, frontend or backend, by active env
func GetActiveEnv(config Config, isBackend bool) string {
	var port = ""
	if isBackend {
		port = "8090"
	} else {
		port = "3000"
	}

	// Quick fix for login redirecting wrong
	if !isBackend && config.Env == "prod" {
		return "https://" + config.Foobar_FrontendHostname
	}

	if config.Env == "prod" {
		return "https://" + config.Foobar_BackendHostname
	} else {
		return "http://localhost:" + port
	}
}

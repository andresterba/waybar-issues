package main

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
)

type Configuration struct {
	Entries []*ConfigurationEntry `json:"entries"`
}

type ConfigurationEntry struct {
	Typ         string `json:"typ"`
	Username    string `json:"username"`
	Token       string `json:"token"`
	URL         string `json:"url"`
	DisplayName string `json:"display_name"`
	OpenIssues  string
}

func getConfigPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/.waybar-issues"
}

func loadConfigFile(filename string, config *Configuration) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	jsonDecoder := json.NewDecoder(file)

	err = jsonDecoder.Decode(&config)
	if err != nil {
		return err
	}

	return nil
}

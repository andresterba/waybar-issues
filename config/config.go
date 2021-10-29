package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/user"
)

var (
	errCouldNotOpenConfigFile  = errors.New("could not open configuration file")
	errCouldNotParseConfigFile = errors.New("could not parse configuration file")
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
}

func GetConfigPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/.waybar-issues"
}

func LoadConfigFile(filename string, config *Configuration) error {
	file, err := os.Open(filename)
	if err != nil {
		return errCouldNotOpenConfigFile
	}

	jsonDecoder := json.NewDecoder(file)

	err = jsonDecoder.Decode(&config)
	if err != nil {
		return errCouldNotParseConfigFile
	}

	return nil
}

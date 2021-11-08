package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/user"
)

type Type string

const (
	GitHubConfigType Type = "github"
	GitLabConfigType Type = "gitlab"
	TrelloConfigType Type = "trello"
)

type Configuration struct {
	Entries []*ConfigurationEntry `json:"entries"`
}

type ConfigurationEntry struct {
	Typ         Type   `json:"typ"`
	Username    string `json:"username"`
	Token       string `json:"token"`
	URL         string `json:"url, omitempty"`
	DisplayName string `json:"display_name"`
}

var (
	errCouldNotOpenConfigFile  = errors.New("could not open configuration file")
	errCouldNotParseConfigFile = errors.New("could not parse configuration file")
)

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

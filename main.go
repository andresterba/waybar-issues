package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/andresterba/waybar-issues/config"
	provider "github.com/andresterba/waybar-issues/provider"
)

type waybarResponse struct {
	Text  string `json:"text"`
	Class string `json:"class"`
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
	configuration := config.Configuration{}

	err := config.LoadConfigFile(config.GetConfigPath(), &configuration)
	checkError(err)

	var providers []provider.Provider

	for _, configEntry := range configuration.Entries {
		switch configEntry.Typ {
		case "gitlab":

			gitLabStatus := provider.NewGitLabStats(configEntry.Username, configEntry.Token, configEntry.URL, configEntry.DisplayName)
			err := gitLabStatus.Process()
			checkError(err)

			providers = append(providers, gitLabStatus)

		case "github":
			gitHubStatus := provider.NewGitHubStats(configEntry.Username, configEntry.Token, configEntry.DisplayName)
			err := gitHubStatus.Process()
			checkError(err)

			providers = append(providers, gitHubStatus)

		default:
			log.Fatal(fmt.Errorf("%s is not supported", configEntry.Typ))
		}
	}

	var responseText string

	for _, provider := range providers {
		responseText += provider.GetFormatedOutput()
	}

	waybarResponse := waybarResponse{Text: responseText, Class: "issues"}
	waybarResponseJSON, err := json.Marshal(waybarResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(waybarResponseJSON))
}

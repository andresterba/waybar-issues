package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/andresterba/waybar-issues/config"
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

	var gitLabInstances []*gitLabStats
	var gitHubInstances []*gitHubStats

	for _, configEntry := range configuration.Entries {
		switch configEntry.Typ {
		case "gitlab":

			gitLabStatus := newGitLabStats(configEntry.Username, configEntry.Token, configEntry.URL, configEntry.DisplayName)
			err := gitLabStatus.process()
			checkError(err)

			gitLabInstances = append(gitLabInstances, gitLabStatus)

		case "github":
			gitHubStatus := newGitHubStats(configEntry.Username, configEntry.Token, configEntry.DisplayName)
			err := gitHubStatus.process()
			checkError(err)

			gitHubInstances = append(gitHubInstances, gitHubStatus)

		default:
			log.Fatal(fmt.Errorf("%s is not supported", configEntry.Typ))
		}
	}

	var responseText string

	for _, gitLabInstance := range gitLabInstances {
		responseText += gitLabInstance.getFormatedOutput()
	}

	for _, gitHubInstance := range gitHubInstances {
		responseText += gitHubInstance.getFormatedOutput()
	}

	waybarResponse := waybarResponse{Text: responseText, Class: "issues"}
	waybarResponseJSON, err := json.Marshal(waybarResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(waybarResponseJSON))
}

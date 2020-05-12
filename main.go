package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/andresterba/waybar-issues/config"

	"github.com/parnurzeal/gorequest"
)

type waybarResponse struct {
	Text  string `json:"text"`
	Class string `json:"class"`
}

func main() {
	configuration := config.Configuration{}

	err := config.LoadConfigFile(config.GetConfigPath(), &configuration)
	if err != nil {
		log.Fatal(err)
	}

	for _, configEntry := range configuration.Entries {
		switch configEntry.Typ {
		case "gitlab":
			openGitLabIssues, err := gitlabGetAssignedIssues(
				configEntry.URL,
				configEntry.Token,
				configEntry.Username,
			)
			if err != nil {
				log.Fatal(err)
			}

			configEntry.OpenIssues = openGitLabIssues

		case "github":
			openGitHubIssues, err := githubGetAssignedIssues(configEntry.Token)
			if err != nil {
				log.Fatal(err)
			}

			configEntry.OpenIssues = openGitHubIssues

		default:
			log.Fatal(fmt.Errorf("%s is not supported", configEntry.Typ))

		}
	}

	var responseText string

	for _, configEntry := range configuration.Entries {
		responseText += configEntry.DisplayName + ": " + configEntry.OpenIssues + " "
	}

	waybarResponse := waybarResponse{Text: responseText, Class: "issues"}
	waybarResponseJSON, err := json.Marshal(waybarResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(waybarResponseJSON))
}

func gitlabGetAssignedIssues(gitlabURL string, gitlabToken string, gitlabUsername string) (string, error) {
	issuesURL := fmt.Sprintf(
		"%s/api/v4/issues?state=opened&scope=assigned_to_me&assignee_username=%s",
		gitlabURL,
		gitlabUsername,
	)

	request := gorequest.New()
	_, body, errs := request.Get(issuesURL).
		AppendHeader("Private-Token", gitlabToken).
		End()
	if errs != nil {
		return "", fmt.Errorf("request to %s failed", issuesURL)
	}

	var issues []interface{}
	if err := json.Unmarshal([]byte(body), &issues); err != nil {
		return "", err
	}
	numberOfIssues := len(issues)

	return strconv.Itoa(numberOfIssues), nil
}

func githubGetAssignedIssues(githubToken string) (string, error) {
	issuesURL := "https://api.github.com/issues"

	request := gorequest.New()
	_, body, errs := request.Get(issuesURL).
		AppendHeader("Authorization", githubToken).
		End()
	if errs != nil {
		return "", fmt.Errorf("request to %s failed", issuesURL)
	}

	var issues []interface{}
	if err := json.Unmarshal([]byte(body), &issues); err != nil {
		return "", err
	}
	numberOfIssues := len(issues)

	return strconv.Itoa(numberOfIssues), nil
}

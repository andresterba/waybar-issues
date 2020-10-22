package main

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

const issueURL = "https://api.github.com/issues"
const pullRequestURL = "https://api.github.com/pulls"

type gitHubStats struct {
	authToken     string
	authUsername  string
	CountOfIssues int
	displayName   string
}

func newGitHubStats(username string, token string, displayName string) *gitHubStats {
	newGitHubStats := &gitHubStats{
		token,
		username,
		0,
		displayName,
	}

	return newGitHubStats
}

func (g *gitHubStats) getAssignedIssues() error {
	request := gorequest.New()
	_, body, errs := request.Get(issueURL).
		AppendHeader("Authorization", g.authToken).
		End()
	if errs != nil {
		return fmt.Errorf("request to %s failed", issueURL)
	}

	var issues []interface{}
	if err := json.Unmarshal([]byte(body), &issues); err != nil {
		return err
	}

	g.CountOfIssues = len(issues)

	return nil
}

func (g *gitHubStats) process() error {
	err := g.getAssignedIssues()
	if err != nil {
		return err
	}

	return nil
}

func (g *gitHubStats) getFormatedOutput() string {
	return fmt.Sprintf("%s: I:%d ", g.displayName, g.CountOfIssues)
}

package provider

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type gitLabStats struct {
	authToken         string
	authUsername      string
	issueCount        int
	mergeRequestCount int
	displayName       string
	instanceURL       string
}

func NewGitLabStats(username string, token string, instanceURL string, displayName string) *gitLabStats {
	newGitLabStats := &gitLabStats{
		authToken:         token,
		authUsername:      username,
		issueCount:        0,
		mergeRequestCount: 0,
		displayName:       displayName,
		instanceURL:       instanceURL,
	}

	return newGitLabStats
}

func (g *gitLabStats) getAssignedIssues() error {
	issuesURL := fmt.Sprintf(
		"%s/api/v4/issues?state=opened&scope=assigned_to_me&assignee_username=%s",
		g.instanceURL,
		g.authUsername,
	)

	request := gorequest.New()
	_, body, errs := request.Get(issuesURL).
		AppendHeader("Private-Token", g.authToken).
		End()
	if errs != nil {
		return fmt.Errorf("request to %s failed", issuesURL)
	}

	var issues []interface{}
	if err := json.Unmarshal([]byte(body), &issues); err != nil {
		return err
	}

	g.issueCount = len(issues)

	return nil
}

func (g *gitLabStats) getAssignedMergeRequests() error {
	mergeRequestURL := fmt.Sprintf(
		"%s/api/v4/merge_requests?state=opened&scope=assigned_to_me&assignee_username=%s",
		g.instanceURL,
		g.authUsername,
	)

	request := gorequest.New()
	_, body, errs := request.Get(mergeRequestURL).
		AppendHeader("Private-Token", g.authToken).
		End()
	if errs != nil {
		return fmt.Errorf("request to %s failed", mergeRequestURL)
	}

	var mergeRequests []interface{}
	if err := json.Unmarshal([]byte(body), &mergeRequests); err != nil {
		return err
	}

	g.mergeRequestCount = len(mergeRequests)

	return nil
}

func (g *gitLabStats) Process() error {
	err := g.getAssignedIssues()
	if err != nil {
		return err
	}

	err = g.getAssignedMergeRequests()
	if err != nil {
		return err
	}

	return nil
}

func (g *gitLabStats) GetFormatedOutput() string {
	return fmt.Sprintf("%s: I:%d MR:%d ", g.displayName, g.issueCount, g.mergeRequestCount)
}

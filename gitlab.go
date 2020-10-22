package main

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type gitLabStats struct {
	authToken         string
	authUsername      string
	IssueCount        int
	MergeRequestCount int
	displayName       string
	instanceURL       string
}

func newGitLabStats(username string, token string, instanceURL string, displayName string) *gitLabStats {
	newGitLabStats := &gitLabStats{
		token,
		username,
		0,
		0,
		displayName,
		instanceURL,
	}

	return newGitLabStats
}

func (g *gitLabStats) gitlabGetAssignedIssues() error {
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

	g.IssueCount = len(issues)

	return nil
}

func (g *gitLabStats) gitlabGetAssignedMergeRequests() error {
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

	g.MergeRequestCount = len(mergeRequests)

	return nil
}

func (g *gitLabStats) process() error {
	err := g.gitlabGetAssignedIssues()
	if err != nil {
		return err
	}

	err = g.gitlabGetAssignedMergeRequests()
	if err != nil {
		return err
	}

	return nil
}

func (g *gitLabStats) getFormatedOutput() string {
	return fmt.Sprintf("%s: I:%d MR:%d ", g.displayName, g.IssueCount, g.MergeRequestCount)
}

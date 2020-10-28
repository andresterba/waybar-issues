package provider

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

const issueURL = "https://api.github.com/issues"

// All of the structs to are used to handle api responses from GitHub are based on
// https://github.com/google/go-github/blob/76c3c3d7c6e78e8c91e77d2e2578c4e0a7cf96ea/github/issues.go#L27
type gitHubStats struct {
	authToken        string
	authUsername     string
	issueCount       int
	pullRequestCount int
	displayName      string
}
type pullRequest struct {
	URL string `json:"url,omitempty"`
}
type issue struct {
	ID          *int64      `json:"id,omitempty"`
	Number      *int        `json:"number,omitempty"`
	Title       *string     `json:"title,omitempty"`
	PullRequest pullRequest `json:"pull_request,omitempty"`
}
type gitHubIssueResponse struct {
	issues []issue `json:""`
}

func NewGitHubStats(username string, token string, displayName string) *gitHubStats {
	newGitHubStats := &gitHubStats{
		token,
		username,
		0,
		0,
		displayName,
	}

	return newGitHubStats
}

func (g *gitHubStats) getAssignedIssuesAndPullRequests() error {
	var issueCounter = 0
	var pullRequestCounter = 0
	var gitHubResponse = new(gitHubIssueResponse)

	request := gorequest.New()
	_, body, errs := request.Get(issueURL).
		AppendHeader("Authorization", g.authToken).
		End()
	if errs != nil {
		return fmt.Errorf("request to %s failed", issueURL)
	}

	if err := json.Unmarshal([]byte(body), &gitHubResponse.issues); err != nil {
		fmt.Println(err)
		return err
	}

	// GitHub's REST API v3 considers every pull request an issue, but not every issue is a pull request.
	// For this reason, "Issues" endpoints may return both issues and pull requests in the response.
	// You can identify pull requests by the pull_request key.
	//
	// https://docs.github.com/en/free-pro-team@latest/rest/reference/issues#list-issues-assigned-to-the-authenticated-user
	for _, issue := range gitHubResponse.issues {
		if len(issue.PullRequest.URL) != 0 {
			pullRequestCounter++
		} else {
			issueCounter++
		}
	}

	g.issueCount = issueCounter
	g.pullRequestCount = pullRequestCounter

	return nil
}

func (g *gitHubStats) Process() error {
	err := g.getAssignedIssuesAndPullRequests()
	if err != nil {
		return err
	}

	return nil
}

func (g *gitHubStats) GetFormatedOutput() string {
	return fmt.Sprintf("%s: I:%d PR:%d ", g.displayName, g.issueCount, g.pullRequestCount)
}

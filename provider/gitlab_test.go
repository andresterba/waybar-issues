package provider

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGitLabProvider(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://my.instance.com/api/v4/merge_requests",
		httpmock.NewStringResponder(200,
			`
			[{
				"name": "MR-1"
			}]
			`))

	httpmock.RegisterResponder("GET", "https://my.instance.com/api/v4/issues",
		httpmock.NewStringResponder(200,
			`
			[{
				"name": "Issue-1"
			},
			{
				"name": "Issue-2"
			}]
			`))

	gitLabStatus := NewGitLabStats("my-test-user", "1337", "https://my.instance.com", "GitLab")
	err := gitLabStatus.Process()
	if err != nil {
		t.Error(err)
	}

	formattedResponse := gitLabStatus.GetFormatedOutput()

	expectedResponse := "GitLab: I:2 MR:1 "

	if formattedResponse != expectedResponse {
		t.Errorf("Response was incorrect, got: %s, want: %s.", formattedResponse, expectedResponse)
	}
}

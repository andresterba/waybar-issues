package provider

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGitHubProviderSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.github.com/issues",
		httpmock.NewStringResponder(200,
			`
			[{
				"id": 1337,
				"number": 123,
				"title": "Test",
				"pull_request": {
					"url": "example.com"
				}
			},
			{
				"id": 1337,
				"number": 123,
				"title": "Test"
			}]
			`))

	gitHubStatus := NewGitHubStats("test", "1337", "GitHub")
	err := gitHubStatus.Process()
	if err != nil {
		t.Error(err)
	}

	formattedResponse := gitHubStatus.GetFormatedOutput()

	expectedResponse := "GitHub: I:1 PR:1 "

	if formattedResponse != expectedResponse {
		t.Errorf("Response was incorrect, got: %s, want: %s.", formattedResponse, expectedResponse)
	}
}

func TestGitHubProviderFailure(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.github.com/issues",
		httpmock.NewStringResponder(404, ""))

	gitHubStatus := NewGitHubStats("test", "1337", "GitHub")
	err := gitHubStatus.Process()
	if err != errAuthenticationFailed {
		t.Errorf("Expected error %s, got: %s.", errAuthenticationFailed, err)
	}
}

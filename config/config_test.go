package config

import (
	"testing"
)

func TestConfigLoader(t *testing.T) {
	configuration := Configuration{}

	err := LoadConfigFile("./config.test", &configuration)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(configuration.Entries) != 2 {
		t.Errorf("Configuration entries size = %d; want 2", len(configuration.Entries))
	}
	if configuration.Entries[0].Typ != "gitlab" {
		t.Errorf("LoadConfigFile() = %s; want gitlab", configuration.Entries[0].Typ)
	}
	if configuration.Entries[0].Username != "test-user" {
		t.Errorf("LoadConfigFile() = %s; want test-user", configuration.Entries[0].Username)
	}
	if configuration.Entries[0].Token != "token1337" {
		t.Errorf("LoadConfigFile() = %s; want token1337", configuration.Entries[0].Token)
	}
	if configuration.Entries[0].URL != "https://gitlab.com" {
		t.Errorf("LoadConfigFile() = %s; want https://gitlab.com", configuration.Entries[0].URL)
	}
	if configuration.Entries[0].DisplayName != "GitLab" {
		t.Errorf("LoadConfigFile() = %s; want GitLab", configuration.Entries[0].DisplayName)
	}
	if configuration.Entries[1].Typ != "github" {
		t.Errorf("LoadConfigFile() = %s; want github", configuration.Entries[1].Typ)
	}
	if configuration.Entries[1].Username != "gh-user" {
		t.Errorf("LoadConfigFile() = %s; want gh-user", configuration.Entries[1].Username)
	}
	if configuration.Entries[1].Token != "token-1337" {
		t.Errorf("LoadConfigFile() = %s; want token-1337", configuration.Entries[1].Token)
	}
	if configuration.Entries[1].URL != "" {
		t.Errorf("LoadConfigFile() = %s; want <empty>", configuration.Entries[1].URL)
	}
	if configuration.Entries[1].DisplayName != "GitHub" {
		t.Errorf("LoadConfigFile() = %s; want GitHub", configuration.Entries[1].DisplayName)
	}
}

func TestConfigOpenError(t *testing.T) {
	configuration := Configuration{}

	err := LoadConfigFile("./no-file.test", &configuration)
	if err != errCouldNotOpenConfigFile {
		t.Errorf("Expected 'could not open file error', got=%T", err)
	}
}

func TestConfigParseError(t *testing.T) {
	configuration := Configuration{}

	err := LoadConfigFile("./invalid_config.test", &configuration)
	if err != errCouldNotParseConfigFile {
		t.Errorf("Expected 'could not parse error', got=%T", err)
	}
}

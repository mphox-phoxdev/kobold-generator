package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFile(t *testing.T) {
	var loadedConfig Config
	err := loadedConfig.LoadConfigurationFile("../test/testconf.json")

	compareConfig := Config{
		LogLevel: "DEBUG",
		Host:     "localhost",
		Port:     "8002",
		Database: Database{
			UserName: "root",
			Password: "password",
			Name:     "kobolds",
			Host:     "localhost",
			Port:     "3306",
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, compareConfig, loadedConfig)
}

func TestLoadConfigFile__FileNotPresent(t *testing.T) {
	var missingConfig Config
	err := missingConfig.LoadConfigurationFile("nofile.json")
	assert.Error(t, err)
}

func TestLoadConfigFile__MalFormedFile(t *testing.T) {
	var malformedConfig Config
	err := malformedConfig.LoadConfigurationFile("../test/malformedconfjson")
	assert.Error(t, err)

}

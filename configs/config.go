package configs

import (
	"encoding/json"
	"os"

	"github.com/juju/errors"
)

type Database struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type Config struct {
	Database     Database `json:"database"`
	LogLevel     string   `json:"logLevel"`
	Host         string   `json:"host"`
	Port         string   `json:"port"`
	CookieDomain string   `json:"cookieDomain"`
	AllowedHosts []string `json:"allowedHosts"`
}

func (config *Config) LoadConfigurationFile(file string) (err error) {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		err = errors.Trace(err)
		return
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		err = errors.Trace(err)
		return
	}

	return
}

package smoke

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	SuiteName string `json:"suite_name"`

	ApiEndpoint string `json:"api"`
	AppsDomain  string `json:"apps_domain"`

	SkipSSLValidation bool `json:"skip_ssl_validation"`

	User string `json:"user"`
	Password string `json:"password"`

	Org string `json:"org"`
	Space string `json:"space"`

	// existing app name - if empty the space will be managed and a random app name will be used
	LoggingApp string `json:"logging_app"`

	ArtifactsDirectory string `json:"artifacts_directory"`
}

// singleton cache
var cachedConfig *Config
func GetConfig() *Config {
	if cachedConfig == nil {
		cachedConfig = loadConfig()
	}
	return cachedConfig
}

func loadConfig() *Config {
	config := newDefaultConfig()
	loadConfigFromJson(config)
	validateRequiredFields(config)
	return config
}

func newDefaultConfig() *Config {
	return &Config{
		ArtifactsDirectory: filepath.Join("..", "results"),
	}
}

func validateRequiredFields(config *Config) {
	if config.SuiteName == "" {
		panic("missing configuration 'suite_name'")
	}

	if config.ApiEndpoint == "" {
		panic("missing configuration 'api'")
	}

	if config.User == "" {
		panic("missing configuration 'user'")
	}

	if config.Password == "" {
		panic("missing configuration 'password'")
	}

	if config.Org == "" {
		panic("missing configuration 'org'")
	}

	if config.Space == "" {
		panic("missing configuration 'space'")
	}
}

// Loads the config from json into the supplied config object
func loadConfigFromJson(config *Config) {
	path := configPath()

	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(config)
	if err != nil {
		panic(err)
	}
}

func configPath() string {
	path := os.Getenv("CONFIG")
	if path == "" {
		panic("Must set $CONFIG to point to an integration config .json file.")
	}

	return path
}

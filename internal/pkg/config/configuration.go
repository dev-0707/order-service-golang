package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Logging  LoggingConfiguration
}

type LoggingConfiguration struct {
	Level string
}

type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Schema       string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	DebugEnabled bool
}

type ServerConfiguration struct {
	Port           int
	Secret         string
	Mode           string
	LogEnabled     bool
	LogFileEnabled bool
	Api            ApiConfiguration
	LogPath        string
}

type ApiConfiguration struct {
	DocsPath       string
	DocsUrl        string
	SwaggerDocsUrl string
}

// SetupDB initialize configuration
func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}

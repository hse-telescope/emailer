package config

import (
	"os"

	"github.com/hse-telescope/utils/queues/kafka"
	"gopkg.in/yaml.v3"
)

// EmailCredentials ...
type EmailCredentials struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

// Config ...
type Config struct {
	Port             uint16                 `yaml:"port"`
	EmailCredentials EmailCredentials       `yaml:"email_credentials"`
	QueueCredentials kafka.QueueCredentials `yaml:"queue_credentials"`
}

func Parse(path string) (Config, error) {
	bytes, err := os.ReadFile(path) // nolint:gosec
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

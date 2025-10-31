package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(path string) (*Config, error) {
	fileData, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config

	err = yaml.Unmarshal(fileData, &cfg)
	if err != nil {
		return nil, fmt.Errorf("parse YAML config: %w", err)
	}

	return &cfg, nil
}

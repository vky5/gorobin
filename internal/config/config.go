package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port    string   `yaml:"port"`
	Servers []string `yaml:"servers"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

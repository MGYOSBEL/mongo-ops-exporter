package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Name string `yaml:"name"`
	URI  string `yaml:"uri"`
}

type Config struct {
	EnableCache  bool             `yaml:"enable-cache"`
	ScrapingTime time.Duration    `yaml:"scraping-time"`
	Databases    []DatabaseConfig `yaml:"databases"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

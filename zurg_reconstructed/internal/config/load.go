package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// LoadZurgConfig loads the configuration from file
func LoadZurgConfig(configPath string) (*ZurgConfig, error) {
	// If no config path provided, try default locations
	if configPath == "" {
		configPath = "config.yml"
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			configPath = filepath.Join("/app", "config.yml")
		}
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file %s: %w", configPath, err)
	}

	var config ZurgConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file %s: %w", configPath, err)
	}

	return &config, nil
}
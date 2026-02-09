package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/EpykLab/wazctl/internal/files"
	"github.com/EpykLab/wazctl/models/configurations"
	"gopkg.in/yaml.v3"
)

var (
	// Lists the defined locations where the config will be located by default
	configLocs = []string{
		".wazctl.yaml",
		"~/.wazctl.yaml",
		"~/.config/wazctl.yaml",
	}
)

// New loads and returns a WazuhCtlConfig from one of the default config locations.
// It returns an error if no valid config file is found or if parsing fails.
func New() (*configurations.WazuhCtlConfig, error) {
	var config configurations.WazuhCtlConfig

	// Iterate through possible config locations
	for _, loc := range configLocs {
		// Expand ~ to home directory
		path, err := expandHomeDir(loc)
		if err != nil {
			log.Printf("Failed to expand path %s: %v", loc, err)
			continue
		}

		// Check if file exists
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}

		// Read the file
		content, err := files.ReadFileFromSpecifiedPath(path)
		if err != nil {
			log.Printf("Failed to read config file at %s: %v", path, err)
			continue
		}

		// Unmarshal YAML into config (pass pointer to update struct)
		if err := yaml.Unmarshal(content, &config); err != nil {
			log.Printf("Failed to unmarshal config file at %s: %v", path, err)
			continue
		}

		// Successfully loaded and parsed config
		return &config, nil
	}

	// No valid config file was found
	return nil, fmt.Errorf("no valid config file found in locations: %v", configLocs)
}

// LoadOptional loads config from the default locations if present.
// If no config file is found, it returns (nil, nil).
// If a file exists but parsing fails, it returns (nil, err).
// Use this when config is optional (e.g. localenv docker) and callers can fall back to defaults.
func LoadOptional() (*configurations.WazuhCtlConfig, error) {
	var config configurations.WazuhCtlConfig

	for _, loc := range configLocs {
		path, err := expandHomeDir(loc)
		if err != nil {
			log.Printf("Failed to expand path %s: %v", loc, err)
			continue
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}

		content, err := files.ReadFileFromSpecifiedPath(path)
		if err != nil {
			log.Printf("Failed to read config file at %s: %v", path, err)
			continue
		}

		if err := yaml.Unmarshal(content, &config); err != nil {
			log.Printf("Failed to unmarshal config file at %s: %v", path, err)
			continue
		}

		return &config, nil
	}

	return nil, nil
}

// expandHomeDir replaces ~ with the user's home directory in the path.
func expandHomeDir(path string) (string, error) {
	if path[:2] != "~/" {
		return path, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(home, path[2:]), nil
}

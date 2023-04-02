// Package config provides functions for caching state
package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

// The base name of gobrief's configuration file
const configFileName = "config.toml"

// configuration allows for saving program configuration and settings
type configuration struct {
	Days      int               `comment:"Days to show in advance"`
	Calendars map[string]string `comment:"List of calendars"`
}

// new returns a default configuration struct
func new() *configuration {
	return &configuration{
		Days:      7,
		Calendars: map[string]string{},
	}
}

// Save writes the configuration to disk
func (s *configuration) Save() error {
	file, err := getSettingsFile()
	if err != nil {
		return err
	}

	dat, err := toml.Marshal(s)
	if err != nil {
		return err
	}

	err = os.WriteFile(file, dat, 0644)
	if err != nil {
		return err
	}

	return nil
}

// getSettingsFile returns the path of mz's cache file.
// The file and its parent directories may not exist.
func getSettingsFile() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil && !os.IsExist(err) {
		return "", err
	}

	folder := filepath.Join(configDir, "gobrief")

	err = os.MkdirAll(folder, 0755)
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(folder, configFileName)

	return configPath, nil
}

// LoadConfig attempts to load the existing configuration file.
// If an error occurs (i.e. the configuration file does not exist)
// a default configuration object is returned.
func LoadConfig() *configuration {
	s := new()

	file, err := getSettingsFile()
	if err != nil {
		log.Println(err)
		return s
	}

	dat, err := os.ReadFile(file)
	if err != nil {
		log.Println(err)
		return s
	}

	err = toml.Unmarshal(dat, &s)
	if err != nil {
		log.Println(err)
		return s
	}

	return s
}

// Package config provides functions for caching state
package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

// Configuration allows for saving program Configuration and settings
type Configuration struct {
	Calendars map[string]string `comment:"List of calendars"`
}

// new returns a default configuration struct
func new() *Configuration {
	return &Configuration{
		Calendars: map[string]string{},
	}
}

// Save writes the configuration to disk
func (s *Configuration) Save() error {
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

// SaveExit is an equivalent to Save(),
// followed by os.Exit(0) if no error occurs.
//
// If an error occurs, log.Fatal is called.
func (s *Configuration) SaveExit() {
	err := s.Save()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
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

	configPath := filepath.Join(folder, "config.toml")

	return configPath, nil
}

// LoadConfig attempts to load the existing configuration file.
// If an error occurs (i.e. the configuration file does not exist)
// a default configuration object is returned.
func LoadConfig() *Configuration {
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

// Package config provides functions for caching state
package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Calendar struct {
	Name     string
	URL      string
	Priority int
}

// Configuration allows for saving program Configuration and settings
type Configuration struct {
	Days      int        `comment:"Upcoming days to show"`
	Calendars []Calendar `comment:"List of calendars"`
}

// new returns a default configuration struct
func new() *Configuration {
	return &Configuration{
		Days:      7,
		Calendars: []Calendar{},
	}
}

// Save writes the configuration to disk
func (s *Configuration) Save() error {
	file, err := openSettingsFile()
	if err != nil {
		return err
	}

	enc := toml.NewEncoder(file)

	return enc.Encode(s)
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

// openSettingsFile returns a file handler to the configuration file.
func openSettingsFile() (*os.File, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	folder := filepath.Join(configDir, "gobrief")

	err = os.MkdirAll(folder, 0755)
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(folder, "config.toml")

	return os.OpenFile(configPath, os.O_RDWR|os.O_CREATE, 0755)
}

// LoadConfig attempts to load the existing configuration file.
// If an error occurs (i.e. the configuration file does not exist)
// a default configuration object is returned.
func LoadConfig() *Configuration {
	s := new()

	file, err := openSettingsFile()
	if err != nil {
		log.Println(err)
		return s
	}

	dec := toml.NewDecoder(file)
	if err != nil {
		log.Println(err)
		return s
	}

	err = dec.Decode(&s)
	return s
}

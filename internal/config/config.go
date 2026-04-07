package config

import (
	"os"
	"encoding/json"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl 				string	`json:"db_url"`
	CurrentUserName		string	`json:"current_user_name"`
}

func Read() (Config, error) {
	// Get Config File Path
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// Open Config File and Defer Closing
	configFile, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()

	// Decode the Content
	decoder := json.NewDecoder(configFile)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(homeDir, configFileName)
	return fullPath, nil	
}

func write(cfg Config) error {
	// Get Config File Path
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Open Config File and Defer Closing
	configFile, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	// Encode Data
	encoder := json.NewEncoder(configFile)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}

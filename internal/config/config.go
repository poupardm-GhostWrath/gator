package config

import (
	"os"
	"encoding/json"
	"fmt"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl 				string	`json:"db_url"`
	CurrentUserName		string	`json:"current_user_name"`
}

func Read() Config {
	// Get Config File Path
	configPath, err := getConfigFilePath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return Config{}
	}

	// Open Config File and Defer Closing
	configFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return Config{}
	}
	defer configFile.Close()

	// Decode the Content
	decoder := json.NewDecoder(configFile)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return Config{}
	}
	return config
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	err := write(*c)
	return err
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + configFileName, nil	
}

func write(cfg Config) error {
	// Get Config File Path
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Open Config File and Defer Closing
	configFile, err := os.OpenFile(configPath, os.O_WRONLY | os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer configFile.Close()

	// Marshal Config struct
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	// Write to File
	_, err = configFile.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

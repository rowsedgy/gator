package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user

	return write(*c)
}

func Read() (Config, error) {
	fileDir, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fileDir)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func write(cfg Config) error {
	fileDir, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fileDir)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(cfg); err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(homeDir, configFileName)
	return fullPath, nil
}

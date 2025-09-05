package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	var conf Config
	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		return Config{}, err
	}
	return conf, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := fmt.Sprintf("%s/%s", home, configFileName)
	return filePath, nil
}

package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Url      string `json:"db_url"`
	Username string `json:"current_user_name,omitempty"`
}

func Read() Config {
	config := Config{}

	configPath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("error getting home directory:", err)
		return Config{}
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return Config{}
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding Json:", err)
		return Config{}
	}
	fmt.Println(config)
	return config
}

func (c Config) SetUser(name string) {
	c.Username = name
	err := write(c)
	if err != nil {
		fmt.Println("there was an error attemping to write user name:", err)
	}
}

func write(c Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("error getting home directory:", err)
		return err
	}

	json, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		fmt.Println("error marshaling Json data.")
		return err
	}

	err = os.WriteFile(configPath, json, 0666)
	if err != nil {
		fmt.Println("error writing Json to file.")
		return err
	}
	fmt.Printf("succesfully wrote username to file: %s\n", c.Username)
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + configFileName, nil
}

package config

import (
	"encoding/json"
	"os"
)

// Configuration is a struct for configuration
// It contains server and database configuration
type Configuration struct {
	Server       *server
	DatabaseDev  *database
	DatabaseTest *database
}

type server struct {
	Host string
	Port string
}

type database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// ParseJSON is used to parse config json file
func ParseJSON(filepath string) (*Configuration, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := &Configuration{}
	err = decoder.Decode(configuration)
	if err != nil {
		return nil, err
	}
	return configuration, nil
}

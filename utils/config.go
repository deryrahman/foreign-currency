package utils

// Configuration is a struct for configuration
// It contains server and database configuration
type Configuration struct {
	Server   *server
	Database *database
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
	return &Configuration{}, nil
}

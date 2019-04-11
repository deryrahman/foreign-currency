package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestParseJSON(t *testing.T) {
	confBody := []byte(`{
		"server": {
			"Host": "123.456.789.0",
			"Port": "8000"
		},
		"database": {
			"Host": "localhost",
			"Port": "1234",
			"User": "dery",
			"Password": "rahman",
			"DBName": "foreigncurrency"
		}
	}`)
	ioutil.WriteFile("./conf.json", confBody, 0644)
	defer os.Remove("./conf.json")

	configuration, _ := ParseJSON("./conf.json")

	assertString(t, configuration.Server.Host, "123.456.789.0")
	assertString(t, configuration.Server.Port, "8000")
	assertString(t, configuration.Database.Host, "localhost")
	assertString(t, configuration.Database.Port, "1234")
	assertString(t, configuration.Database.User, "dery")
	assertString(t, configuration.Database.Password, "rahman")
	assertString(t, configuration.Database.DBName, "foreigncurrency")
}

func TestParseJSON_fail(t *testing.T) {
	configuration, err := ParseJSON("./conf.json")
	if configuration != nil {
		t.Errorf("should return nil")
	}
	if err == nil {
		t.Errorf("wanted an error")
	}
}

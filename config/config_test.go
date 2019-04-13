package config

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

func assertError(t *testing.T, err error, wantMsg string) {
	if err == nil {
		t.Errorf("wanted an error")
	}
	assertString(t, err.Error(), wantMsg)
}

func TestParseJSON(t *testing.T) {
	confBody := []byte(`{
		"Server": {
			"Host": "123.456.789.0",
			"Port": "8000"
		},
		"DatabaseDev": {
			"Host": "localhost",
			"Port": "1234",
			"User": "dery",
			"Password": "rahman",
			"DBName": "foreigncurrency"
		},
		"DatabaseTest": {
			"Host": "localhost_test",
			"Port": "1235",
			"User": "dery_test",
			"Password": "rahman_test",
			"DBName": "test"
		}
	}`)
	ioutil.WriteFile("./conf.json", confBody, 0644)
	defer os.Remove("./conf.json")

	configuration, _ := ParseJSON("./conf.json")

	assertString(t, configuration.Server.Host, "123.456.789.0")
	assertString(t, configuration.Server.Port, "8000")
	assertString(t, configuration.DatabaseDev.Host, "localhost")
	assertString(t, configuration.DatabaseDev.Port, "1234")
	assertString(t, configuration.DatabaseDev.User, "dery")
	assertString(t, configuration.DatabaseDev.Password, "rahman")
	assertString(t, configuration.DatabaseDev.DBName, "foreigncurrency")
	assertString(t, configuration.DatabaseTest.Host, "localhost_test")
	assertString(t, configuration.DatabaseTest.Port, "1235")
	assertString(t, configuration.DatabaseTest.User, "dery_test")
	assertString(t, configuration.DatabaseTest.Password, "rahman_test")
	assertString(t, configuration.DatabaseTest.DBName, "test")
}

func TestParseJSON_failNofile(t *testing.T) {
	configuration, err := ParseJSON("./conf.json")
	if configuration != nil {
		t.Errorf("should return nil")
	}
	if err == nil {
		t.Errorf("wanted an error")
	}
	assertError(t, err, "open ./conf.json: no such file or directory")
}

func TestParseJSON_failInvalidParse(t *testing.T) {
	confBody := []byte(`{
		"SS": {"host": "localhost}
	}`)
	ioutil.WriteFile("./conf.json", confBody, 0644)
	defer os.Remove("./conf.json")
	configuration, err := ParseJSON("./conf.json")
	if configuration != nil {
		t.Errorf("should return nil")
	}
	if err == nil {
		t.Errorf("wanted an error")
	}
	assertError(t, err, "invalid character '\\n' in string literal")
}

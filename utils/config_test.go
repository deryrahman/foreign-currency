package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

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

	want := "123.456.789.0"
	got := configuration.Server.Host
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
	want = "8000"
	got = configuration.Server.Port
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
	want = "localhost"
	got = configuration.Database.Host
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
	want = "1234"
	got = configuration.Database.Port
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
	want = "dery"
	got = configuration.Database.User
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
	want = "rahman"
	got = configuration.Database.Password
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
	want = "foreigncurrency"
	got = configuration.Database.DBName
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

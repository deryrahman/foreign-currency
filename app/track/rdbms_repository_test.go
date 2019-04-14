package track

import (
	"fmt"
	"log"
	"testing"

	"github.com/deryrahman/foreign-currency/app"
	"github.com/deryrahman/foreign-currency/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func newDB(t *testing.T) *gorm.DB {
	configuration, err := config.ParseJSON("../../config.json")
	if err != nil {
		log.Fatalf("couldn't parse config. %s\n", err.Error())
	}
	database := configuration.DatabaseTest
	dsl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", database.User, database.Password, database.Host, database.Port, database.DBName)
	db, err := gorm.Open("mysql", dsl)
	if err != nil {
		t.Fatal(err.Error())
	}
	db.DropTableIfExists(&app.Rate{}, &app.Currency{}, &app.Track{})
	db.AutoMigrate(&app.Rate{}, &app.Currency{}, &app.Track{})
	return db
}

func assertUint(t *testing.T, got, want uint) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func assertInt(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
func TestFetch(t *testing.T) {
	db := newDB(t)
	defer db.Close()
	currencies := []app.Currency{
		app.Currency{
			From: "SGD",
			To:   "USD",
		},
		app.Currency{
			From: "IDR",
			To:   "USD",
		},
		app.Currency{
			From: "JPY",
			To:   "USD",
		},
	}
	for i := range currencies {
		db.Create(&currencies[i])
	}
	tracks := []app.Track{
		app.Track{
			CurrencyID: currencies[0].ID,
		},
		app.Track{
			CurrencyID: currencies[1].ID,
		},
	}
	for i := range tracks {
		db.Create(&tracks[i])
	}

	repo := CreateRDBMSRepo(db)
	gots, _ := repo.Fetch()
	assertInt(t, len(gots), 2)
	for i := range gots {
		assertUint(t, gots[i].ID, uint(i+1))
		assertUint(t, gots[i].CurrencyID, currencies[i].ID)
	}
}

func TestStore(t *testing.T) {
	db := newDB(t)
	defer db.Close()

	currencies := []app.Currency{
		app.Currency{
			From: "SGD",
			To:   "USD",
		},
		app.Currency{
			From: "IDR",
			To:   "USD",
		},
		app.Currency{
			From: "JPY",
			To:   "USD",
		},
	}
	for i := range currencies {
		db.Create(&currencies[i])
	}
	tracks := []app.Track{
		app.Track{
			CurrencyID: currencies[0].ID,
		},
		app.Track{
			CurrencyID: currencies[1].ID,
		},
	}

	repo := CreateRDBMSRepo(db)
	for i := range tracks {
		repo.Store(&tracks[i])
	}
	gots := []app.Track{}
	db.Find(&gots)
	assertInt(t, len(gots), len(tracks))
	for i := range gots {
		assertUint(t, gots[i].ID, uint(i+1))
		assertUint(t, gots[i].CurrencyID, tracks[i].CurrencyID)
	}
}

func TestStore_exist(t *testing.T) {
	db := newDB(t)
	defer db.Close()

	currencies := []app.Currency{
		app.Currency{
			From: "SGD",
			To:   "USD",
		},
		app.Currency{
			From: "IDR",
			To:   "USD",
		},
		app.Currency{
			From: "JPY",
			To:   "USD",
		},
	}
	for i := range currencies {
		db.Create(&currencies[i])
	}
	tracks := []app.Track{
		app.Track{
			CurrencyID: currencies[0].ID,
		},
		app.Track{
			CurrencyID: currencies[1].ID,
		},
	}
	for i := range tracks {
		db.Create(&tracks[i])
	}

	repo := CreateRDBMSRepo(db)
	for i := range tracks {
		tracks[i].ID = 0
		err := repo.Store(&tracks[i])
		if err == nil {
			t.Errorf("wanted an error")
		}
	}
}

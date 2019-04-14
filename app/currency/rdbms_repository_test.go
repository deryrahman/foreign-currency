package currency

import (
	"fmt"
	"log"
	"testing"
	"time"

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
	dsl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", database.User, database.Password, database.Host, database.Port, database.DBName)
	db, err := gorm.Open("mysql", dsl)
	if err != nil {
		t.Fatal(err.Error())
	}
	db.DropTableIfExists(&app.Rate{}, &app.Currency{})
	db.AutoMigrate(&app.Rate{}, &app.Currency{})
	return db
}

func assertBool(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
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
		app.Currency{From: "USD", To: "SGD"},
	}
	db.Create(&currencies[0])

	repo := CreateRDBMSRepo(db)
	gots, _ := repo.Fetch()
	assertUint(t, gots[0].ID, 1)
	assertString(t, gots[0].From, currencies[0].From)
	assertString(t, gots[0].To, currencies[0].To)
}

func TestFetchOne_fullRates(t *testing.T) {
	db := newDB(t)
	defer db.Close()

	currencies := []app.Currency{
		app.Currency{From: "USD", To: "SGD"},
	}
	ti := time.Now()
	rates := []app.Rate{
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
	}
	db.Create(&currencies[0])
	for i := range rates {
		db.Create(&rates[i])
	}
	repo := CreateRDBMSRepo(db)
	got, _ := repo.FetchOne("USD", "SGD", -1)
	assertUint(t, got.ID, 1)
	assertString(t, got.From, currencies[0].From)
	assertString(t, got.To, currencies[0].To)
	assertInt(t, len(got.Rates), 11)
}
func TestFetchOne_partialRates(t *testing.T) {
	db := newDB(t)
	defer db.Close()

	currencies := []app.Currency{
		app.Currency{From: "USD", To: "SGD"},
	}
	ti := time.Now()
	rates := []app.Rate{
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
	}
	db.Create(&currencies[0])
	for i := range rates {
		db.Create(&rates[i])
	}
	repo := CreateRDBMSRepo(db)
	got, _ := repo.FetchOne("USD", "SGD", 3)
	assertUint(t, got.ID, 1)
	assertString(t, got.From, currencies[0].From)
	assertString(t, got.To, currencies[0].To)
	assertInt(t, len(got.Rates), 3)
}

func TestFetchOne_zeroRates(t *testing.T) {
	db := newDB(t)
	defer db.Close()

	currencies := []app.Currency{
		app.Currency{From: "USD", To: "SGD"},
	}
	ti := time.Now()
	rates := []app.Rate{
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
	}
	db.Create(&currencies[0])
	for i := range rates {
		db.Create(&rates[i])
	}
	repo := CreateRDBMSRepo(db)
	got, _ := repo.FetchOne("USD", "SGD", 0)
	assertUint(t, got.ID, 1)
	assertString(t, got.From, currencies[0].From)
	assertString(t, got.To, currencies[0].To)
	assertInt(t, len(got.Rates), 0)
}

func TestUpdate(t *testing.T) {
	db := newDB(t)
	defer db.Close()

	currencies := []app.Currency{
		app.Currency{From: "USD", To: "SGD", Tracked: false, TrackedRev: false},
	}
	db.Create(&currencies[0])
	repo := CreateRDBMSRepo(db)

	currencies[0].Tracked = true
	got, _ := repo.Update(1, &currencies[0])
	assertUint(t, got.ID, 1)
	assertBool(t, got.Tracked, true)
}

func TestStore(t *testing.T) {
	db := newDB(t)
	defer db.Close()

	currencies := []app.Currency{
		app.Currency{From: "USD", To: "SGD"},
	}
	repo := CreateRDBMSRepo(db)
	repo.Store(&currencies[0])

	got := app.Currency{}
	db.First(&got, "currencies.from = ? AND currencies.to = ?", "USD", "SGD")

	assertUint(t, got.ID, 1)
	assertString(t, got.From, currencies[0].From)
	assertString(t, got.To, currencies[0].To)
}

func TestStore_exist(t *testing.T) {
	db := newDB(t)
	defer db.Close()

	currencies := []app.Currency{
		app.Currency{From: "USD", To: "SGD"},
	}
	db.Create(&currencies[0])
	currencies[0].ID = 0

	repo := CreateRDBMSRepo(db)
	err := repo.Store(&currencies[0])

	if err == nil {
		t.Errorf("wanted an error")
	}
}

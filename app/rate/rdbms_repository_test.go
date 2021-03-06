package rate

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
	dsl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", database.User, database.Password, database.Host, database.Port, database.DBName)
	db, err := gorm.Open("mysql", dsl)
	if err != nil {
		t.Fatal(err.Error())
	}
	db.DropTableIfExists(&app.Rate{}, &app.Currency{})
	db.AutoMigrate(&app.Rate{}, &app.Currency{})
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
	gots, _ := repo.Fetch()
	assertInt(t, len(gots), 11)
	for i := range gots {
		assertUint(t, gots[i].ID, uint(i+1))
	}
}

func TestFetchBetweenDate(t *testing.T) {
	db := newDB(t)
	defer db.Close()
	currencies := []app.Currency{
		app.Currency{From: "USD", To: "SGD"},
	}

	const RFC3339FullDate = "2006-01-02"
	rates := []app.Rate{}
	db.Create(&currencies[0])
	for i := 0; i < 11; i++ {
		ti, _ := time.Parse(RFC3339FullDate, fmt.Sprintf("2019-08-%02d", 12+i))
		rate := app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1}
		rates = append(rates, rate)
		db.Create(&rates[i])
	}

	repo := CreateRDBMSRepo(db)
	from, _ := time.Parse(RFC3339FullDate, "2019-08-01")
	to, _ := time.Parse(RFC3339FullDate, "2019-08-15")
	gots, _ := repo.FetchBetweenDate(1, &from, &to)
	assertInt(t, len(gots), 4)
	for i := range gots {
		assertUint(t, gots[i].ID, uint(i+1))
	}
}

func TestStore(t *testing.T) {
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
	repo := CreateRDBMSRepo(db)

	for i := range rates {
		repo.Store(&rates[i])
	}

	gots := []app.Rate{}
	db.Find(&gots, "rates.currency_id = ?", 1)
	for i := range rates {
		assertUint(t, gots[i].ID, uint(i+1))
	}
}

func TestStore_exist(t *testing.T) {
	db := newDB(t)
	defer db.Close()

	currencies := []app.Currency{
		app.Currency{From: "USD", To: "SGD"},
	}
	const RFC3339FullDate = "2006-01-02"
	ti, _ := time.Parse(RFC3339FullDate, "2019-08-12")

	rates := []app.Rate{
		app.Rate{Date: &ti, RateValue: 0.6, CurrencyID: 1},
	}
	db.Create(&currencies[0])
	db.Create(&rates[0])
	rates[0].ID = 0

	repo := CreateRDBMSRepo(db)
	err := repo.Store(&rates[0])
	if err == nil {
		t.Errorf("wanted an error")
	}
}

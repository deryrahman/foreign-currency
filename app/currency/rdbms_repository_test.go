package currency

import (
	"regexp"
	"testing"
	"time"

	"github.com/deryrahman/foreign-currency/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func newDB() (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("mysql", db)
	return mock, gormDB
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
	mock, db := newDB()
	defer db.Close()
	currencies := []app.Currency{
		app.Currency{ID: 1, From: "USD", To: "SGD"},
	}
	rows := mock.NewRows([]string{"id", "from", "to"}).
		AddRow(currencies[0].ID, currencies[0].From, currencies[0].To)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `currencies`")).WillReturnRows(rows)

	repo := CreateRDBMSRepo(db)
	gots, _ := repo.Fetch()
	assertUint(t, gots[0].ID, currencies[0].ID)
	assertString(t, gots[0].From, currencies[0].From)
	assertString(t, gots[0].To, currencies[0].To)
}

func TestFetchOne_full(t *testing.T) {
	mock, db := newDB()
	defer db.Close()
	currencies := []app.Currency{
		app.Currency{ID: 1, From: "USD", To: "SGD"},
	}
	ti := time.Now()
	rates := []app.Rate{
		app.Rate{ID: 1, Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{ID: 2, Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{ID: 3, Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{ID: 4, Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{ID: 5, Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{ID: 6, Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{ID: 7, Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{ID: 8, Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{ID: 9, Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{ID: 10, Date: &ti, RateValue: 0.6, CurrencyID: 1},
		app.Rate{ID: 11, Date: &ti, RateValue: 0.6, CurrencyID: 1},
	}
	rowsCurrencies := mock.NewRows([]string{"id", "from", "to"}).
		AddRow(currencies[0].ID, currencies[0].From, currencies[0].To)
	rowsRates := mock.NewRows([]string{"id", "date", "rate_value", "currency_id"})
	for i := range rates {
		rowsRates = rowsRates.AddRow(rates[i].ID, rates[i].Date, rates[i].RateValue, rates[i].CurrencyID)
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `currencies` WHERE (from = ? AND to = ?) ORDER BY `currencies`.`id` ASC LIMIT 1")).WillReturnRows(rowsCurrencies)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `rates`  WHERE (`currency_id` = ?) ORDER BY rates.date DESC,`rates`.`id` ASC")).WillReturnRows(rowsRates)

	repo := CreateRDBMSRepo(db)
	got, _ := repo.FetchOne("USD", "SGD", -1)
	assertUint(t, got.ID, currencies[0].ID)
	assertString(t, got.From, currencies[0].From)
	assertString(t, got.To, currencies[0].To)
	assertInt(t, len(got.Rates), 11)
}

package currency

import (
	"regexp"
	"testing"

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

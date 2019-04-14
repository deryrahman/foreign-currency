package track

import (
	"github.com/deryrahman/foreign-currency/app"
	"github.com/jinzhu/gorm"
)

// RDBMSRepo is a struct to wrap its DB
type RDBMSRepo struct {
	DB *gorm.DB
}

// CreateRDBMSRepo is used to create new RDBMS repository
func CreateRDBMSRepo(db *gorm.DB) *RDBMSRepo {
	return &RDBMSRepo{db}
}

// Fetch is a method to fetch all list of currency track
func (repo *RDBMSRepo) Fetch() ([]*app.Track, error) {
	return nil, nil
}

// Store is a method to store currency rate to track
func (repo *RDBMSRepo) Store(*app.Track) error {
	return nil
}

// DeleteByID is method to delete a track by its ID
func (repo *RDBMSRepo) DeleteByID(uint) (*app.Track, error) {
	return nil, nil
}

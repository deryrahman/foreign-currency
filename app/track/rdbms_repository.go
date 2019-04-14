package track

import (
	"errors"

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
	tracks := []app.Track{}
	repo.DB.Find(&tracks)
	result := []*app.Track{}
	for i := range tracks {
		result = append(result, &tracks[i])
	}
	return result, nil
}

// Store is a method to store currency rate to track
// If there's existing track (same currencyID and Revert), this method will throw error
func (repo *RDBMSRepo) Store(track *app.Track) error {
	repo.DB.First(track, "tracks.currency_id = ? AND tracks.revert = ?", track.CurrencyID, track.Revert)
	ok := repo.DB.NewRecord(track)
	if !ok {
		return errors.New("track exist")
	}
	repo.DB.Create(track)
	return nil
}

// DeleteByID is method to delete a track by its ID
func (repo *RDBMSRepo) DeleteByID(uint) (*app.Track, error) {
	return nil, nil
}

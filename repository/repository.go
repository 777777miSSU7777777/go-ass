package repository

import (
	"fmt"

	"github.com/777777miSSU7777777/go-ass/model"

	"github.com/jinzhu/gorm"
)

type UploadTrackCallback func(id int64) error

type DeleteTrackCallback func(id int64) error

type Repository struct {
	db *gorm.DB
}

func NewRepository(dbType string, connectionString string) *Repository {
	db, err := gorm.Open(dbType, connectionString)
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	db.AutoMigrate(&model.Artist{})
	db.AutoMigrate(&model.Genre{})
	db.AutoMigrate(&model.GenreTracks{})
	db.AutoMigrate(&model.Playlist{})
	db.AutoMigrate(&model.PlaylistTracks{})
	db.AutoMigrate(&model.Track{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserTokens{})
	db.AutoMigrate(&model.UserTracks{})

	fmt.Println("Database migrated")

	return &Repository{ db: db }
}

func (repo *Repository) GetAllArtists() ([]model.Artist, error) {
	var artists []model.Artist
	if err := repo.db.Find(&artists).Error; err != nil {
		return nil, err
	}
	
	return artists, nil
}

func (repo *Repository) GetArtist(artistID int64) (model.Artist, error) {
	var artist model.Artist
	if err := repo.db.where("artist_id", artistID).First(&artist).Error; if err != nil {
		return model.Artist, err
	}

	return artist
}

func (repo *Repository) AddNewArtist(newArtist model.Artist) (model.Artist, err) {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollBack()
		}
	}()

	result := tx.Create(&newArtist)

	if err := result.Error; err != nil {
		tx.Rollback()
		return model.Artist{}, err
	}

	resultValue := result.Value.(model.Artist)

	if err := tx.Commit().Error; err != nil {
		tx.RollBack()
		return model.Artist{}, err
	}

	return resultValue, nil
}

func (repo *Repository) UpdateArtist(updatedArtist model.Artist) (model.Artist, error) {
	var artist model.Artist

	if err := repo.db.Where("artist_id", updatedArtist.ArtistID).Find(&artist).Error; err != nil {
		return model.Artist{}, err
	}

	if err := repo.db.Model(&artist).Updates(updatedArtist).Error; err != nil {
		return model.Artist{}, err
	}

	return artist, nil
}

func (repo *Repository) DeleteArtist(artistID int64) error {
	if err := repo.db.Where("artist_id", artistID).Delete(&model.Artist{}).Error; err != nil {
		return err
	}

	return nil
}

func (repo *Repository) GetAllTracks() ([]model.Track, error) {
	var tracks []model.Track
	if err := repo.db.Find(&tracks).Error; err != nil {
		return nil, err
	}

	return tracks, nil
}

func (repo *Repository) GetTrack(trackID int64) (model.Track, error) {
	var track model.Track
	if err := repo.db.Where("track_id", trackID).First(&track).Error; if err != nil {
		return model.Track{}, err
	}

	return track
}

func (repo *Repository) AddNewTrack(newTrack model.Track, uploadTrack UploadTrackCallback) (model.Track, err) {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollBack()
		}
	}()

	result := tx.Create(&newTrack);

	if err := result.Error; err != nil {
		tx.Rollback();
		return model.Track{}, err
	}

	resultValue := result.Value.(model.Track)

	if err := uploadTrack(); err != nil {
		tx.Rollback();
		return model.Track{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback();
		return model.Track{}, err
	}

	return resultValue, nil
}

func (repo *Repository) UpdateTrack(updatedTrack model.Track) (model.Track, error) {
	var track model.Track

	if err := repo.db.Where("track_id", updatedTrack.TrackID).Find(&track).Error; err != nil {
		return model.Track{}, err
	}

	if err := repo.db.Model(&track).Updates(updatedTrack).Error; err != nil {
		return model.Track{}, err
	}

	return track, nil
}

func (repo *Repository) DeleteTrack(trackID int64) error {
	if err := repo.db.Where("track_id", trackID).Delete(&model.Track{}).Error; err != nil {
		return err
	}

	return nil
}
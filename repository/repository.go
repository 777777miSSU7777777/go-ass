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
	db.AutoMigrate(&model.UserPlaylists{})
	db.AutoMigrate(&model.UserTokens{})
	db.AutoMigrate(&model.UserTracks{})

	fmt.Println("Database migrated")

	return &Repository{db: db}
}

func (repo *Repository) GetAllArtists() ([]model.Artist, error) {
	var artists []model.Artist
	err := repo.db.Find(&artists).Error
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (repo *Repository) GetArtist(artistID int64) (model.Artist, error) {
	var artist model.Artist
	err := repo.db.Where(&model.Artist{ArtistID: artistID}).First(&artist).Error
	if err != nil {
		return model.Artist{}, err
	}

	return artist, nil
}

func (repo *Repository) AddNewArtist(newArtist model.Artist) (model.Artist, error) {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Create(&newArtist)

	err := result.Error
	if err != nil {
		tx.Rollback()
		return model.Artist{}, err
	}

	resultValue := result.Value.(model.Artist)

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return model.Artist{}, err
	}

	return resultValue, nil
}

func (repo *Repository) UpdateArtist(updatedArtist model.Artist) (model.Artist, error) {
	var artist model.Artist

	err := repo.db.Where(model.Artist{ArtistID: updatedArtist.ArtistID}).Find(&artist).Error
	if err != nil {
		return model.Artist{}, err
	}

	err = repo.db.Model(&artist).Updates(updatedArtist).Error
	if err != nil {
		return model.Artist{}, err
	}

	return artist, nil
}

func (repo *Repository) DeleteArtist(artistID int64) error {
	err := repo.db.Delete(&model.Artist{ArtistID: artistID}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) GetAllGenres() ([]model.Genre, error) {
	var genres []model.Genre
	err := repo.db.Find(&genres).Error
	if err != nil {
		return nil, err
	}

	return genres, nil
}

func (repo *Repository) GetGenre(genreID int64) (model.Genre, error) {
	var genre model.Genre
	err := repo.db.Where(&model.Genre{GenreID: genreID}).First(&genre).Error
	if err != nil {
		return model.Genre{}, err
	}

	return genre, nil
}

func (repo *Repository) AddNewGenre(newGenre model.Genre) (model.Genre, error) {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Create(&newGenre)

	err := result.Error
	if err != nil {
		tx.Rollback()
		return model.Genre{}, err
	}

	resultValue := result.Value.(model.Genre)

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return model.Genre{}, err
	}

	return resultValue, nil
}

func (repo *Repository) UpdateGenre(updatedGenre model.Genre) (model.Genre, error) {
	var genre model.Genre

	err := repo.db.Where(&model.Genre{GenreID: updatedGenre.GenreID}).Find(&genre).Error
	if err != nil {
		return model.Genre{}, err
	}

	err = repo.db.Model(&genre).Updates(updatedGenre).Error
	if err != nil {
		return model.Genre{}, err
	}

	return genre, nil
}

func (repo *Repository) GetAllPlaylists(updatedPlaylist model.Playlist) ([]model.Playlist, error) {
	var playlists []model.Playlist
	err := repo.db.Find(&playlists).Error
	if err != nil {
		return nil, err
	}

	return playlists, nil
}

func (repo *Repository) GetPlaylist(playlistID int64) (model.Playlist, error) {
	var playlist model.Playlist
	err := repo.db.Where(&model.Playlist{PlaylistID: playlistID}).First(&playlist).Error
	if err != nil {
		return model.Playlist{}, err
	}

	return playlist, nil
}

func (repo *Repository) AddNewPlaylist(newPlaylist model.Playlist) (model.Playlist, error) {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Create(&newPlaylist)

	err := result.Error
	if err != nil {
		tx.Rollback()
		return model.Playlist{}, err
	}

	resultValue := result.Value.(model.Playlist)

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return model.Playlist{}, err
	}

	return resultValue, nil
}

func (repo *Repository) UpdatePlaylist(updatedPlaylist model.Playlist) (model.Playlist, error) {
	var playlist model.Playlist

	err := repo.db.Where(&model.Playlist{PlaylistID: updatedPlaylist.PlaylistID}).Find(&playlist).Error
	if err != nil {
		return model.Playlist{}, err
	}

	err = repo.db.Model(&playlist).Updates(updatedPlaylist).Error
	if err != nil {
		return model.Playlist{}, err
	}

	return playlist, nil
}

func (repo *Repository) DeletePlaylist(playlistID int64) error {
	err := repo.db.Delete(&model.Playlist{PlaylistID: playlistID}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) AddTracksToPlaylist(playlistID int64, tracksID ...int64) error {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	for _, trackID := range tracksID {
		playlistTracks := model.PlaylistTracks{PlaylistID: playlistID, TrackID: trackID}

		err := repo.db.Create(&playlistTracks).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *Repository) DeleteTracksFromPlaylist(playlistID int64, tracksID ...int64) error {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	for _, trackID := range tracksID {
		err := repo.db.Delete(&model.PlaylistTracks{PlaylistID: playlistID, TrackID: trackID}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *Repository) GetAllTracks() ([]model.Track, error) {
	var tracks []model.Track
	err := repo.db.Find(&tracks).Error
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (repo *Repository) GetTrack(trackID int64) (model.Track, error) {
	var track model.Track
	err := repo.db.Where(&model.Track{TrackID: trackID}).First(&track).Error
	if err != nil {
		return model.Track{}, err
	}

	return track, nil
}

func (repo *Repository) AddNewTrack(newTrack model.Track, uploadTrack UploadTrackCallback) (model.Track, error) {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Create(&newTrack)

	err := result.Error
	if err != nil {
		tx.Rollback()
		return model.Track{}, err
	}

	resultValue := result.Value.(model.Track)

	if resultValue.GenreID != 0 {
		genreTracks := model.GenreTracks{GenreID: resultValue.GenreID, TrackID: resultValue.TrackID}
		err := tx.Create(&genreTracks).Error
		if err != nil {
			tx.Rollback()
			return model.Track{}, err
		}
	}

	err = uploadTrack(resultValue.TrackID)
	if err != nil {
		tx.Rollback()
		return model.Track{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return model.Track{}, err
	}

	return resultValue, nil
}

func (repo *Repository) UpdateTrack(updatedTrack model.Track) (model.Track, error) {
	var track model.Track

	err := repo.db.Where(&model.Track{TrackID: updatedTrack.TrackID}).Find(&track).Error
	if err != nil {
		return model.Track{}, err
	}

	err = repo.db.Model(&track).Updates(updatedTrack).Error
	if err != nil {
		return model.Track{}, err
	}

	return track, nil
}

func (repo *Repository) DeleteTrack(trackID int64) error {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	err := repo.db.Delete(&model.Track{TrackID: trackID}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = repo.db.Delete(&model.GenreTracks{TrackID: trackID}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = repo.db.Delete(&model.UserTracks{TrackID: trackID}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = repo.db.Delete(&model.PlaylistTracks{TrackID: trackID}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *Repository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *Repository) GetUser(userID int64) (model.User, error) {
	var user model.User
	err := repo.db.Where(&model.User{UserID: userID}).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *Repository) AddNewUser(newUser model.User) (model.User, error) {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Create(&newUser)

	err := result.Error
	if err != nil {
		tx.Rollback()
		return model.User{}, nil
	}

	resultValue := result.Value.(model.User)

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	return resultValue, nil
}

func (repo *Repository) UpdateUser(updatedUser model.User) (model.User, error) {
	var user model.User

	err := repo.db.Where(&model.User{UserID: updatedUser.UserID}).Find(&user).Error
	if err != nil {
		return model.User{}, err
	}

	err = repo.db.Model(&user).Updates(updatedUser).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *Repository) DeleteUser(userID int64) error {
	err := repo.db.Delete(&model.User{UserID: userID}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) AddTracksToUserList(userID int64, tracksID ...int64) error {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	for _, trackID := range tracksID {
		userTracks := model.UserTracks{UserID: userID, TrackID: trackID}

		err := repo.db.Create(&userTracks).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *Repository) DeleteTracksFromUserList(userID int64, tracksID ...int64) error {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	for _, trackID := range tracksID {
		err := repo.db.Delete(&model.UserTracks{UserID: userID, TrackID: trackID}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *Repository) AddPlaylistsToUserList(userID int64, playlistsID ...int64) error {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	for _, playlistID := range playlistsID {
		userPlaylists := model.UserPlaylists{UserID: userID, PlaylistID: playlistID}

		err := repo.db.Create(&userPlaylists).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *Repository) DeletePlaylistsFromUserList(userID int64, playlistsID ...int64) error {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	for _, playlistID := range playlistsID {
		err := repo.db.Delete(&model.UserPlaylists{UserID: userID, PlaylistID: playlistID}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

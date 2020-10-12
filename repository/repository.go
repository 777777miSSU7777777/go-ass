package repository

import (
	"fmt"

	"github.com/777777miSSU7777777/go-ass/helper"
	"github.com/777777miSSU7777777/go-ass/model"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(connectionString string) Repository {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
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

	return Repository{db: db}
}

func (repo *Repository) GetAllArtists() ([]model.Artist, error) {
	var artists []model.Artist
	err := repo.db.Find(&artists).Error
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (repo *Repository) GetArtist(artistID string) (model.Artist, error) {
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

	newArtist.ArtistID = uuid.New().String()
	result := tx.Create(&newArtist)

	err := result.Error
	if err != nil {
		tx.Rollback()
		return model.Artist{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return model.Artist{}, err
	}

	return newArtist, nil
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

func (repo *Repository) DeleteArtist(artistID string) error {
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

func (repo *Repository) GetGenre(genreID string) (model.Genre, error) {
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

	newGenre.GenreID = uuid.New().String()
	result := tx.Create(&newGenre)

	err := result.Error
	if err != nil {
		tx.Rollback()
		return model.Genre{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return model.Genre{}, err
	}

	return newGenre, nil
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

func (repo *Repository) GetAllPlaylists() ([]model.Playlist, error) {
	var playlists []model.Playlist
	err := repo.db.Find(&playlists).Error
	if err != nil {
		return nil, err
	}

	return playlists, nil
}

func (repo *Repository) GetUserPlaylists(userID string) ([]model.UserPlaylists, error) {
	var userPlaylists []model.UserPlaylists
	err := repo.db.Where(model.UserPlaylists{UserID: userID}).Find(&userPlaylists).Error
	if err != nil {
		return nil, err
	}

	return userPlaylists, nil
}

func (repo *Repository) GetPlaylist(playlistID string) (model.Playlist, error) {
	var playlist model.Playlist
	err := repo.db.Where(&model.Playlist{PlaylistID: playlistID}).First(&playlist).Error
	if err != nil {
		return model.Playlist{}, err
	}

	return playlist, nil
}

func (repo *Repository) GetPlaylistTracks(playlistID string) ([]model.PlaylistTracks, error) {
	var playlistTracks []model.PlaylistTracks
	err := repo.db.Where(&model.PlaylistTracks{PlaylistID: playlistID}).Find(&playlistTracks).Error
	if err != nil {
		return nil, err
	}

	return playlistTracks, nil
}

func (repo *Repository) AddNewPlaylist(newPlaylist model.Playlist) (model.Playlist, error) {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	newPlaylist.PlaylistID = uuid.New().String()
	result := tx.Create(&newPlaylist)

	err := result.Error
	if err != nil {
		tx.Rollback()
		return model.Playlist{}, err
	}

	userPlaylists := model.UserPlaylists{UserID: newPlaylist.CreatedByID, PlaylistID: newPlaylist.PlaylistID}
	err = tx.Create(&userPlaylists).Error
	if err != nil {
		tx.Rollback()
		return model.Playlist{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return model.Playlist{}, err
	}

	return newPlaylist, nil
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

func (repo *Repository) DeletePlaylist(playlistID string) error {
	err := repo.db.Delete(&model.Playlist{PlaylistID: playlistID}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) AddTracksToPlaylist(playlistID string, tracksID ...string) error {
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

func (repo *Repository) DeleteTracksFromPlaylist(playlistID string, tracksID ...string) error {
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

func (repo *Repository) GetTrack(trackID string) (model.Track, error) {
	var track model.Track
	err := repo.db.Where(&model.Track{TrackID: trackID}).First(&track).Error
	if err != nil {
		return model.Track{}, err
	}

	return track, nil
}

func (repo *Repository) AddNewTrack(newTrack model.Track, uploadTrack helper.UploadTrackCallback) (model.Track, error) {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	newTrack.TrackID = uuid.New().String()
	result := tx.Create(&newTrack)

	err := result.Error
	if err != nil {
		tx.Rollback()
		return model.Track{}, err
	}

	if newTrack.ArtistID != "" {
		var artist model.Artist
		err := tx.Where(&model.Artist{ArtistID: newTrack.ArtistID}).First(&artist).Error
		if err != nil {
			tx.Rollback()
			return model.Track{}, err
		}
	}

	if newTrack.GenreID != "" {
		var genre model.Genre
		err := tx.Where(&model.Genre{GenreID: newTrack.GenreID}).First(&genre).Error
		if err != nil {
			tx.Rollback()
			return model.Track{}, err
		}

		genreTracks := model.GenreTracks{GenreID: newTrack.GenreID, TrackID: newTrack.TrackID}
		err = tx.Create(&genreTracks).Error
		if err != nil {
			tx.Rollback()
			return model.Track{}, err
		}
	}

	if newTrack.UploadedByID != "" {
		userTracks := model.UserTracks{UserID: newTrack.UploadedByID, TrackID: newTrack.TrackID}
		err := tx.Create(&userTracks).Error
		if err != nil {
			tx.Rollback()
			return model.Track{}, err
		}
	}

	err = uploadTrack(newTrack.TrackID)
	if err != nil {
		tx.Rollback()
		return model.Track{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return model.Track{}, err
	}

	return newTrack, nil
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

func (repo *Repository) DeleteTrack(trackID string, deleteTrack helper.DeleteTrackCallback) error {
	tx := repo.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Delete(&model.Track{TrackID: trackID}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Delete(&model.GenreTracks{TrackID: trackID}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Delete(&model.UserTracks{TrackID: trackID}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Delete(&model.PlaylistTracks{TrackID: trackID}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = deleteTrack()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
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

func (repo *Repository) GetUser(userID string) (model.User, error) {
	var user model.User
	err := repo.db.Where(&model.User{UserID: userID}).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *Repository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := repo.db.Where(&model.User{Email: email}).First(&user).Error
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

	newUser.UserID = uuid.New().String()
	result := tx.Create(&newUser)

	err := result.Error
	if err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	return newUser, nil
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

func (repo *Repository) DeleteUser(userID string) error {
	err := repo.db.Delete(&model.User{UserID: userID}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) AddRefreshToken(userID string, refreshToken string) error {
	userTokens := model.UserTokens{UserID: userID, Token: refreshToken}
	err := repo.db.Create(&userTokens).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UpdateRefreshToken(userID string, refreshToken string, newRefreshToken string) error {
	var userTokens model.UserTokens
	newUserTokens := model.UserTokens{UserID: userID, Token: newRefreshToken}
	err := repo.db.Where(&model.UserTokens{UserID: userID, Token: refreshToken}).Find(&userTokens).Error
	if err != nil {
		return err
	}

	err = repo.db.Model(&userTokens).Updates(newUserTokens).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) DeleteRefreshToken(userID string, refreshToken string) error {
	err := repo.db.Delete(&model.UserTokens{UserID: userID, Token: refreshToken}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) GetUserTrackList(userID string) ([]model.UserTracks, error) {
	var userTracks []model.UserTracks

	err := repo.db.Where(&model.UserTracks{UserID: userID}).Find(&userTracks).Error
	if err != nil {
		return nil, err
	}

	return userTracks, nil
}

func (repo *Repository) AddTracksToUserList(userID string, tracksID ...string) error {
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

func (repo *Repository) DeleteTracksFromUserList(userID string, tracksID ...string) error {
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

func (repo *Repository) AddPlaylistsToUserList(userID string, playlistsID ...string) error {
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

func (repo *Repository) DeletePlaylistsFromUserList(userID string, playlistsID ...string) error {
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

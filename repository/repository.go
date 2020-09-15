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
	if err := repo.db.Where(&model.Artist{ ArtistID: artistID }).First(&artist).Error; if err != nil {
		return model.Artist{}, err
	}

	return artist, nil
}

func (repo *Repository) AddNewArtist(newArtist model.Artist) (model.Artist, error) {
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

	if err := repo.db.Where(model.Artist{ ArtistID: updatedArtist.ArtistID }).Find(&artist).Error; err != nil {
		return model.Artist{}, err
	}

	if err := repo.db.Model(&artist).Updates(updatedArtist).Error; err != nil {
		return model.Artist{}, err
	}

	return artist, nil
}

func (repo *Repository) DeleteArtist(artistID int64) error {
	if err := repo.db.Delete(&model.Artist{ ArtistID: artistID }).Error; err != nil {
		return err
	}

	return nil
}

func (repo *Repository) GetAllGenres() ([]model.Genre, error) {
	var genres []model.Genre
	if err := repo.db.Find(&genres).Error; err != nil {
		return nil, err
	}

	return genres, nil
}

func (repo *Repository) GetGenre(genreID int64) (model.Genre, error) {
	var genre model.Genre
	if err := repo.db.Where(&model.Genre{ GenreID: genreID }).First(&genre).Error; if err != nil {
		return model.Genre{}, err
	}

	return genre, nil
}

func (repo *Repository) AddNewGenre(newGenre model.Genre) (model.Genre, error) {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollBack()
		}
	}()

	result := tx.Create(&newGenre)

	if err := result.Error; err != nil {
		tx.Rollback()
		return model.Genre{}, err
	}

	resultValue := result.Value.(model.Genre)

	if err := tx.Commit().Error; err != nil {
		tx.RollBack()
		return model.Genre{}, err
	}

	return resultValue, nil
}

func (repo *Repository) UpdateGenre(updatedGenre model.Genre) (model.Genre, error) {
	var genre model.Genre

	if err := repo.db.Where(&model.Genre{ GenreID: updatedGenre.GenreID }).Find(&genre).Error; err != nil {
		return model.Genre{}, err
	}

	if err := repo.db.Model(&genre).Updates(updatedGenre).Error; err != nil {
		return model.Genre{}, err
	}

	return genre, nil
}

func (repo *Repository) GetAllPlaylists(updatedPlaylist model.Playlist) (model.Playlist, error) {
	var playlists []model.Playlist
	if err := repo.db.Find(&playlists).Error; err != nil {
		return nil, err
	}

	return playlists, nil
}

func (repo *Repository) GetPlaylist(playlistID int64) (model.Playlist, error) {
	var playlist model.Playlist
	if err := repo.db.Where(&model.Playlist{ PlaylistID: playlistID }).First(&playlist).Error; if err != nil {
		return model.Playlist{}, err
	}

	return playlist, nil
}

func (repo *Repository) AddNewPlaylist(newPlaylist model.Playlist) (model.Playlist, error) {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollBack()
		}
	}()

	result := tx.Create(&newPlaylist)

	if err := result.Error; err != nil {
		tx.RollBack()
		return model.Playlist{}, err
	}

	resultValue := result.Value.(model.Playlist)

	if err := tx.Commit().Error; err != nil {
		tx.RollBack()
		return model.Playlist{}, err
	}

	return resultValue, nil
}

func (repo *Repository) UpdatePlaylist(updatedPlaylist model.Playlist) (model.Playlist, error) {
	var playlist model.Playlist

	if err := repo.db.Where(&model.Playlist{ PlaylistID: updatedPlaylist.PlaylistID }).Find(&playlist).Error; err != nil {
		return model.Playlist{}, err
	}

	if err := repo.db.Model(&playlist).Updates(updatedPlaylist).Error; err != nil {
		return model.Playlist{}, err
	}

	return playlist, nil
}

func (repo *Repository) DeletePlaylist(playlistID int64) error {
	if err := repo.db.Delete(&model.Playlist{ PlaylistID: playlistID }).Error; err != nil {
		return err
	}

	return nil
}

func (repo *Repository) AddTracksToPlaylist(playlistID int64, tracksID ...int64) error {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollBack()
		}
	}()

	for _, trackID := range tracksID {
		playlistTracks := model.PlaylistTracks{ playlistID: playlistID, trackID: trackID }

		err := repo.db.Create(&playlistTracks).Error; if err != nil {
			tx.RollBack()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.RollBack()
		return err
	}

	return nil
}

func (repo *Repository) DeleteTracksFromPlaylist(playlistID int64, tracksID ...int64) error {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollBack()
		}
	}()

	for _, trackID := range tracksID {
		err := repo.db.Delete(&model.PlaylistTracks{ PlaylistID: playlistID, TrackID: trackID }).Error; if err != nil {
			tx.RollBack()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.RollBack()
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
	if err := repo.db.Where(&model.Track{ TrackID: trackID }).First(&track).Error; if err != nil {
		return model.Track{}, err
	}

	return track, nil
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

	if (resultValue.GenreID) {
		genreTracks := model.GenreTracks{ GenreID: resultValue.GenreID, TrackID: resultValue.TrackID }
		if err := tx.Create(&genreTracks).Error; if err != nil {
			tx.RollBack();
			return model.Track{}, err
		}
	}

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

	if err := repo.db.Where(&model.Track{ TrackID: updatedTrack.TrackID }).Find(&track).Error; err != nil {
		return model.Track{}, err
	}

	if err := repo.db.Model(&track).Updates(updatedTrack).Error; err != nil {
		return model.Track{}, err
	}

	return track, nil
}

func (repo *Repository) DeleteTrack(trackID int64) error {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollBack()
		}
	}()

	if err := repo.db.Delete(&model.Track{ TrackID: trackID }).Error; err != nil {
		tx.RollBack()
		return err
	}

	if err = repo.db.Delete(&model.GenreTrack{ TrackID: trackID }).Error; err != nil {
		tx.RollBack()
		return err
	}

	if err = repo.db.Delete(&model.UserTracks{ TrackID: trackID }).Error; err != nil {
		tx.RollBack()
		return err
	}

	if err = repo.db.Delete(&model.PlaylistTracks{ TrackID: trackID }).Error; err != nil {
		tx.RollBack()
		return err
	}

	return nil
}

func (repo *Repository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	
	return users, nil
}

func (repo *Repository) GetUser(userID int64) (model.User, error) {
	var user model.User
	if err := repo.db.Where(&model.User{ UserID: userID }).First(&user).Error; if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *Repository) AddNewUser(newUser model.User) (model.User, error) {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollBack()
		}
	}()

	result := tx.Create(&newUser)

	if err := result.Error; err != nil {
		tx.RollBack()
		return model.User{}, nil
	}

	resultValue := result.Value.(model.User)

	if err := tx.Commit().Error; err != nil {
		tx.RollBack()
		return model.User{}, err
	}

	return resultValue, nil
}

func (repo *Repository) UpdateUser(updatedUser model.User) (model.User, error) {
	var user model.User

	if err := repo.db.Where(&model.User{ UserID: updatedUser.UserID }).Find(&user).Error; err != nil {
		return model.Artist{}, err
	}

	if err := repo.db.Model(&user).Updates(updatedUser).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *Repository) DeleteUser(userID int64) (model.User, error) {
	if err := repo.db.Delete(&model.User{ UserID: userID }).Error; err != nil {
		return err
	}

	return nil
}

func (repo *Repository) AddTracksToUserList(userID int64, tracksID ...int64) error {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollBack()
		}
	}()


	for _, trackID := range tracksID {
		userTracks := model.UserTracks{ userID: userID, trackID: trackID }
	
		err := repo.db.Create(&userTracks).Error; if err != nil {
			tx.RollBack()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.RollBack()
		return err
	}

	return nil
}

func (repo *Repository) DeleteTracksFromUserList(userID int64, tracksID ...) {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollBack()
		}
	}()

	for _, trackID := range tracksID {
		err := repo.db.Delete(&model.UserTracks{ UserID: userID, TrackID: trackID }).Error; if err != nil {
			tx.RollBack()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.RollBack()
		return err
	}

	return nil
}

func (repo *Repository) AddPlaylistsToUserList(userID int64, playlistsID ...int64) error {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollBack()
		}
	}()

	for _, playlistID := range playlistsID {
		userPlaylists := model.UserPlaylists{ userID: userID, playlistID: playlistID }

		err := repo.db.Create(&userPlaylists).Error; if err != nil {
			tx.RollBack()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.RollBack()
		return err
	}

	return nil
}

func (repo *Repository) DeletePlaylistsFromUserList(userID int64, playlistsID ...int64) error {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollBack()
		}
	}()

	for _, playlistID := range playlistsID {
		err := repo.db.Delete(&model.UserPlaylists{ UserID: userID, playlistID: PlaylistID }).Error; if err != nil {
			tx.RollBack()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.RollBack()
		return err
	}

	return nil
}
package service

import (
	"fmt"

	"github.com/777777miSSU7777777/go-ass/helper"
	"github.com/777777miSSU7777777/go-ass/model"
	"github.com/777777miSSU7777777/go-ass/repository"
)

var UserCredentialsAreInvalidError = fmt.Errorf("user credentials are invalid error")

type Service struct {
	repo repository.Repository
}

func New(r repository.Repository) Service {
	return Service{r}
}

func (service Service) AddTrack(title string, artistID int64, genreID int64, uploadedByID int64, uploadTrack helper.UploadTrackCallback) (model.Track, error) {
	newTrack := model.Track{TrackTitle: title, ArtistID: artistID, GenreID: genreID, UploadedByID: uploadedByID}
	track, err := service.repo.AddNewTrack(newTrack, uploadTrack)
	if err != nil {
		return model.Track{}, err
	}

	return track, nil
}

func (service Service) GetAllTracks() ([]model.Track, error) {
	tracks, err := service.repo.GetAllTracks()
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (service Service) GetTrackByID(trackID int64) (model.Track, error) {
	track, err := service.repo.GetTrack(trackID)
	if err != nil {
		return model.Track{}, err
	}

	return track, nil
}

func (service Service) UpdateTrackByID(trackID int64, title string, artistID int64, genreID int64) (model.Track, error) {
	updatedTrack := model.Track{TrackID: trackID, TrackTitle: title, ArtistID: artistID, GenreID: genreID}
	track, err := service.repo.UpdateTrack(updatedTrack)
	if err != nil {
		return model.Track{}, err
	}

	return track, nil
}

func (service Service) DeleteTrackByID(trackID int64, deleteTrack helper.DeleteTrackCallback) error {
	err := service.repo.DeleteTrack(trackID, deleteTrack)
	if err != nil {
		return err
	}

	return nil
}

func (service Service) SignUp(email string, username string, password string) error {
	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return err
	}

	newUser := model.User{Email: email, Username: username, Password: hashedPassword, Role: helper.UserRole}
	_, err = service.repo.AddNewUser(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (service Service) SignIn(email string, password string) (string, string, error) {
	user, err := service.repo.GetUserByEmail(email)
	if err != nil {
		return "", "", err
	}

	if helper.CheckPasswordHash(password, user.Password) {
		accessToken, refreshToken, err := helper.GenerateTokens(user.UserID)
		if err != nil {
			return "", "", err
		}

		return accessToken, refreshToken, nil
	}

	return "", "", nil
}

func (service Service) RefreshToken(token string) (string, string, error) {
	tokenClaims, err := helper.ParseToken(token)

	if err != nil {
		return "", "", err
	}

	userID := tokenClaims["userID"].(int64)

	accessToken, refreshToken, err := helper.GenerateTokens(userID)
	if err != nil {
		return "", "", err
	}

	err = service.repo.UpdateRefreshToken(userID, token, refreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (service Service) SignOut(token string) error {
	tokenClaims, err := helper.ParseToken(token)
	if err != nil {
		return err
	}

	userID := tokenClaims["userID"].(int64)

	err = service.repo.DeleteRefreshToken(userID, token)

	if err != nil {
		return err
	}

	return nil
}

func (service Service) GetUserTrackList(userID int64) ([]model.Track, error) {
	return nil, nil
}

func (service Service) AddTrackToUserTrackList(userID int64, trackID int64) ([]model.Track, error) {
	return nil, nil
}

func (service Service) RemoveTrackFromUserTrackList(userID int64, trackID int64) ([]model.Track, error) {
	return nil, nil
}

func (service Service) GetAllPlaylists() ([]model.Playlist, [][]model.Track, error) {
	return nil, nil, nil
}

func (service Service) GetUserPlaylists(userID int64) ([]model.Playlist, [][]model.Track, error) {
	return nil, nil, nil
}

func (service Service) CreateNewPlaylist(title string, createdByID int64, trackList []int64) (model.Playlist, []model.Track, error) {
	return model.Playlist{}, nil, nil
}

func (service Service) GetPlaylistByID(playlistID int64) (model.Playlist, []model.Track, error) {
	return model.Playlist{}, nil, nil
}

func (service Service) DeletePlaylistByID(playlistID int64, createdByID int64) error {
	return nil
}

func (service Service) AddTracksToPlaylist(userID int64, playlistID int64, trackList []string) (model.Playlist, []model.Track, error) {
	return model.Playlist{}, nil, nil
}

func (service Service) RemoveTracksFromPlaylist(userID int64, playlistID int64, trackList []string) (model.Playlist, []model.Track, error) {
	return model.Playlist{}, nil, nil
}

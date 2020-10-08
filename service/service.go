package service

import (
	"fmt"

	"github.com/777777miSSU7777777/go-ass/helper"
	"github.com/777777miSSU7777777/go-ass/model"
	"github.com/777777miSSU7777777/go-ass/repository"
)

var UserCredentialsAreInvalidError = fmt.Errorf("user credentials are invalid error")
var InvalidUserError = fmt.Errorf("user id is invalid error")

type Service struct {
	repo repository.Repository
}

func New(r repository.Repository) Service {
	return Service{r}
}

func (service Service) AddNewTrack(title string, artistID int64, genreID int64, uploadedByID int64, uploadTrack helper.UploadTrackCallback) (model.TrackResponse, error) {
	newTrack := model.Track{TrackTitle: title, ArtistID: artistID, GenreID: genreID, UploadedByID: uploadedByID}
	dbTrack, err := service.repo.AddNewTrack(newTrack, uploadTrack)
	if err != nil {
		return model.TrackResponse{}, err
	}

	trackArtist, err := service.repo.GetArtist(dbTrack.ArtistID)
	if err != nil {
		return model.TrackResponse{}, err
	}

	trackResponse := model.TrackResponse{ID: dbTrack.TrackID, Title: dbTrack.TrackTitle, Artist: trackArtist.ArtistName}
	return trackResponse, nil
}

func (service Service) GetAllTracks() ([]model.TrackResponse, error) {
	dbTracks, err := service.repo.GetAllTracks()
	if err != nil {
		return nil, err
	}

	var responseTracks []model.TrackResponse = make([]model.TrackResponse, 0, len(dbTracks))

	for _, dbTrack := range dbTracks {
		trackArtist, err := service.repo.GetArtist(dbTrack.ArtistID)
		if err != nil {
			return nil, err
		}

		trackResponse := model.TrackResponse{ID: dbTrack.TrackID, Title: dbTrack.TrackTitle, Artist: trackArtist.ArtistName}
		responseTracks = append(responseTracks, trackResponse)
	}

	return responseTracks, nil
}

func (service Service) GetTrackByID(trackID int64) (model.TrackResponse, error) {
	dbTrack, err := service.repo.GetTrack(trackID)
	if err != nil {
		return model.TrackResponse{}, err
	}

	trackArtist, err := service.repo.GetArtist(dbTrack.ArtistID)
	if err != nil {
		return model.TrackResponse{}, err
	}

	trackResponse := model.TrackResponse{ID: dbTrack.TrackID, Title: dbTrack.TrackTitle, Artist: trackArtist.ArtistName}
	return trackResponse, nil
}

func (service Service) UpdateTrackByID(trackID int64, title string, artistID int64, genreID int64) (model.TrackResponse, error) {
	updatedTrack := model.Track{TrackID: trackID, TrackTitle: title, ArtistID: artistID, GenreID: genreID}
	dbTrack, err := service.repo.UpdateTrack(updatedTrack)
	if err != nil {
		return model.TrackResponse{}, err
	}

	trackArtist, err := service.repo.GetArtist(dbTrack.ArtistID)
	if err != nil {
		return model.TrackResponse{}, err
	}

	trackResponse := model.TrackResponse{ID: dbTrack.TrackID, Title: dbTrack.TrackTitle, Artist: trackArtist.ArtistName}
	return trackResponse, nil
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

func (service Service) GetUserTrackList(userID int64) ([]model.TrackResponse, error) {
	userTracks, err := service.repo.GetUserTrackList(userID)
	if err != nil {
		return nil, err
	}

	dbTracks := make([]model.Track, 0, len(userTracks))
	for _, userTrack := range userTracks {
		dbTrack, err := service.repo.GetTrack(userTrack.TrackID)
		if err != nil {
			return nil, err
		}

		dbTracks = append(dbTracks, dbTrack)
	}

	responseTracks := make([]model.TrackResponse, 0, len(dbTracks))

	for _, dbTrack := range dbTracks {
		trackArtist, err := service.repo.GetArtist(dbTrack.ArtistID)
		if err != nil {
			return nil, err
		}

		trackResponse := model.TrackResponse{ID: dbTrack.TrackID, Title: dbTrack.TrackTitle, Artist: trackArtist.ArtistName}
		responseTracks = append(responseTracks, trackResponse)
	}

	return responseTracks, nil
}

func (service Service) AddTracksToUserTrackList(userID int64, tracksID ...int64) error {
	err := service.repo.AddTracksToUserList(userID, tracksID...)
	if err != nil {
		return err
	}

	return nil
}

func (service Service) DeleteTracksFromUserTrackList(userID int64, tracksID ...int64) error {
	err := service.repo.DeleteTracksFromUserList(userID, tracksID...)
	if err != nil {
		return err
	}

	return nil
}

func (service Service) GetAllPlaylists() ([]model.PlaylistResponse, error) {
	dbPlaylists, err := service.repo.GetAllPlaylists()
	if err != nil {
		return nil, err
	}

	responsePlaylists := make([]model.PlaylistResponse, 0, len(dbPlaylists))
	for _, dbPlaylist := range dbPlaylists {
		dbPlaylistTracks, err := service.repo.GetPlaylistTracks(dbPlaylist.PlaylistID)
		if err != nil {
			return nil, err
		}

		playlistTracksResponse := make([]model.TrackResponse, 0, len(dbPlaylistTracks))
		for _, dbPlaylistTrack := range dbPlaylistTracks {
			dbTrack, err := service.repo.GetTrack(dbPlaylistTrack.TrackID)
			if err != nil {
				return nil, err
			}

			trackArtist, err := service.repo.GetArtist(dbTrack.ArtistID)
			if err != nil {
				return nil, err
			}

			trackResponse := model.TrackResponse{ID: dbTrack.TrackID, Title: dbTrack.TrackTitle, Artist: trackArtist.ArtistName}
			playlistTracksResponse = append(playlistTracksResponse, trackResponse)
		}

		playlistResponse := model.PlaylistResponse{ID: dbPlaylist.PlaylistID, Title: dbPlaylist.PlaylistTitle, TrackList: playlistTracksResponse}
		responsePlaylists = append(responsePlaylists, playlistResponse)
	}

	return responsePlaylists, nil
}

func (service Service) GetUserPlaylists(userID int64) ([]model.PlaylistResponse, error) {
	dbUserPlaylists, err := service.repo.GetUserPlaylists(userID)
	if err != nil {
		return nil, err
	}

	dbPlaylists := make([]model.Playlist, 0, len(dbUserPlaylists))

	for _, dbUserPlaylist := range dbUserPlaylists {
		dbPlaylist, err := service.repo.GetPlaylist(dbUserPlaylist.PlaylistID)
		if err != nil {
			return nil, err
		}

		dbPlaylists = append(dbPlaylists, dbPlaylist)
	}

	responsePlaylists := make([]model.PlaylistResponse, 0, len(dbPlaylists))
	for _, dbPlaylist := range dbPlaylists {
		dbPlaylistTracks, err := service.repo.GetPlaylistTracks(dbPlaylist.PlaylistID)
		if err != nil {
			return nil, err
		}

		playlistTracksResponse := make([]model.TrackResponse, 0, len(dbPlaylistTracks))
		for _, dbPlaylistTrack := range dbPlaylistTracks {
			dbTrack, err := service.repo.GetTrack(dbPlaylistTrack.TrackID)
			if err != nil {
				return nil, err
			}

			trackArtist, err := service.repo.GetArtist(dbTrack.ArtistID)
			if err != nil {
				return nil, err
			}

			trackResponse := model.TrackResponse{ID: dbTrack.TrackID, Title: dbTrack.TrackTitle, Artist: trackArtist.ArtistName}
			playlistTracksResponse = append(playlistTracksResponse, trackResponse)
		}

		playlistResponse := model.PlaylistResponse{ID: dbPlaylist.PlaylistID, Title: dbPlaylist.PlaylistTitle, TrackList: playlistTracksResponse}
		responsePlaylists = append(responsePlaylists, playlistResponse)
	}

	return responsePlaylists, nil
}

func (service Service) CreateNewPlaylist(title string, createdByID int64, trackList []int64) (model.PlaylistResponse, error) {
	newPlaylist := model.Playlist{PlaylistTitle: title, CreatedByID: createdByID}
	dbPlaylist, err := service.repo.AddNewPlaylist(newPlaylist)
	if err != nil {
		return model.PlaylistResponse{}, err
	}

	err = service.repo.AddTracksToPlaylist(dbPlaylist.PlaylistID, trackList...)
	if err != nil {
		return model.PlaylistResponse{}, err
	}

	dbPlaylistTracks, err := service.repo.GetPlaylistTracks(dbPlaylist.PlaylistID)
	if err != nil {
		return model.PlaylistResponse{}, err
	}

	playlistTracksResponse := make([]model.TrackResponse, 0, len(dbPlaylistTracks))
	for _, dbPlaylistTrack := range dbPlaylistTracks {
		dbTrack, err := service.repo.GetTrack(dbPlaylistTrack.TrackID)
		if err != nil {
			return model.PlaylistResponse{}, err
		}

		trackArtist, err := service.repo.GetArtist(dbTrack.ArtistID)
		if err != nil {
			return model.PlaylistResponse{}, err
		}

		trackResponse := model.TrackResponse{ID: dbTrack.TrackID, Title: dbTrack.TrackTitle, Artist: trackArtist.ArtistName}
		playlistTracksResponse = append(playlistTracksResponse, trackResponse)
	}

	playlistResponse := model.PlaylistResponse{ID: dbPlaylist.PlaylistID, Title: dbPlaylist.PlaylistTitle, TrackList: playlistTracksResponse}

	return playlistResponse, nil
}

func (service Service) GetPlaylistByID(playlistID int64) (model.PlaylistResponse, error) {
	dbPlaylist, err := service.repo.GetPlaylist(playlistID)
	if err != nil {
		return model.PlaylistResponse{}, err
	}

	dbPlaylistTracks, err := service.repo.GetPlaylistTracks(playlistID)
	if err != nil {
		return model.PlaylistResponse{}, err
	}

	playlistTracksResponse := make([]model.TrackResponse, 0, len(dbPlaylistTracks))
	for _, dbPlaylistTrack := range dbPlaylistTracks {
		dbTrack, err := service.repo.GetTrack(dbPlaylistTrack.TrackID)
		if err != nil {
			return model.PlaylistResponse{}, err
		}

		trackArtist, err := service.repo.GetArtist(dbTrack.ArtistID)
		if err != nil {
			return model.PlaylistResponse{}, err
		}

		trackResponse := model.TrackResponse{ID: dbTrack.TrackID, Title: dbTrack.TrackTitle, Artist: trackArtist.ArtistName}
		playlistTracksResponse = append(playlistTracksResponse, trackResponse)
	}

	playlistResponse := model.PlaylistResponse{ID: dbPlaylist.PlaylistID, Title: dbPlaylist.PlaylistTitle, TrackList: playlistTracksResponse}

	return playlistResponse, nil
}

func (service Service) DeletePlaylistByID(playlistID int64, userID int64) error {
	dbPlaylist, err := service.repo.GetPlaylist(playlistID)
	if err != nil {
		return err
	}

	if dbPlaylist.CreatedByID != userID {
		return InvalidUserError
	}

	err = service.repo.DeletePlaylist(playlistID)
	if err != nil {
		return err
	}

	return nil
}

func (service Service) AddTracksToPlaylist(userID int64, playlistID int64, trackList []int64) error {
	dbPlaylist, err := service.repo.GetPlaylist(playlistID)
	if err != nil {
		return err
	}

	if dbPlaylist.CreatedByID != userID {
		return InvalidUserError
	}

	err = service.repo.AddTracksToPlaylist(playlistID, trackList...)
	if err != nil {
		return err
	}

	return nil
}

func (service Service) DeleteTracksFromPlaylist(userID int64, playlistID int64, trackList []int64) error {
	dbPlaylist, err := service.repo.GetPlaylist(playlistID)
	if err != nil {
		return err
	}

	if dbPlaylist.CreatedByID != userID {
		return InvalidUserError
	}

	err = service.repo.DeleteTracksFromPlaylist(playlistID, trackList...)
	if err != nil {
		return err
	}

	return nil
}

func (service Service) AddPlaylistsToUserList(userID int64, playlistsID ...int64) error {
	err := service.repo.AddPlaylistsToUserList(userID, playlistsID...)
	if err != nil {
		return err
	}

	return nil
}

func (service Service) DeletePlaylistsFromUserList(userID int64, playlistsID ...int64) error {
	err := service.repo.DeletePlaylistsFromUserList(userID, playlistsID...)
	if err != nil {
		return err
	}

	return nil
}

package service

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

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

func (s Service) AddTrack(author, title, uploadedByID string) (model.Track, error) {
	err := model.ValidateTrack(author, title)
	if err != nil {
		return model.Track{}, err
	}

	id, err := s.repo.AddTrack(author, title, uploadedByID)
	if err != nil {
		return model.Track{}, err
	}

	audio, err := s.repo.GetTrackByID(id)
	if err != nil {
		return model.Track{}, err
	}

	return audio, nil
}

func (s Service) GetAllTracks() ([]model.Track, error) {
	tracks, err := s.repo.GetAllTracks()
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (s Service) GetTrackByID(trackID string) (model.Track, error) {
	track, err := s.repo.GetTrackByID(trackID)
	if err != nil {
		return model.Track{}, err
	}

	return track, nil
}

func (s Service) GetTracksByKey(key string) ([]model.Track, error) {
	tracks, err := s.repo.GetTracksByKey(key)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (s Service) UpdateTrackByID(trackID string, author, title string) (model.Track, error) {
	err := model.ValidateTrack(author, title)
	if err != nil {
		return model.Track{}, err
	}

	err = s.repo.UpdateTrackByID(trackID, author, title)
	if err != nil {
		return model.Track{}, err
	}

	track, err := s.repo.GetTrackByID(trackID)
	if err != nil {
		return model.Track{}, err
	}

	return track, nil
}

func (s Service) DeleteTrackByID(trackID string) error {
	err := s.repo.DeleteTrackByID(trackID)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) SignUp(email, name, password string) error {
	err := model.ValidateUser(email, name, password)
	if err != nil {
		return err
	}

	_, err = s.repo.AddUser(email, name, password)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) SignIn(email, password string) (string, string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return "", "", UserCredentialsAreInvalidError
		}
		return "", "", err
	} else {
		customClaims := repository.JWTPayload{
			user.ID.Hex(),
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Second * time.Duration(1800)).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
		accessToken, err := token.SignedString([]byte(repository.SecretKey))
		if err != nil {
			return "", "", fmt.Errorf("error while signing user refresh token: %v", err)
		}

		refreshToken, err := s.repo.AddRefreshToken(user.ID.Hex())
		if err != nil {
			return "", "", err
		}

		return accessToken, refreshToken, nil
	}
}

func (s Service) RefreshToken(token string) (string, string, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &repository.JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(repository.SecretKey), nil
	})

	if err != nil {
		return "", "", fmt.Errorf("error while parsing refresh token: %v", err)
	}

	payload := jwtToken.Claims.(*repository.JWTPayload)

	customClaims := repository.JWTPayload{
		payload.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(1800)).Unix(),
		},
	}
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	accessToken, err := unsignedToken.SignedString([]byte(repository.SecretKey))
	if err != nil {
		return "", "", fmt.Errorf("error while signing user access token: %v", err)
	}

	refreshToken, err := s.repo.UpdateRefreshToken(token)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s Service) SignOut(token string) error {
	err := s.repo.DeleteRefreshToken(token)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetUserTrackList(userID string) ([]model.Track, error) {
	userTrackList, err := s.repo.GetUserTrackList(userID)
	if err != nil {
		return nil, err
	}

	return userTrackList, nil
}

func (s Service) AddTrackToUserTrackList(userID, trackID string) ([]model.Track, error) {
	err := s.repo.AddTrackToUserTrackList(userID, trackID)
	if err != nil {
		return nil, err
	}

	userTrackList, err := s.repo.GetUserTrackList(userID)
	if err != nil {
		return nil, err
	}

	return userTrackList, nil
}

func (s Service) RemoveTrackFromUserTrackList(userID, trackID string) ([]model.Track, error) {
	err := s.repo.RemoveTrackFromUserTrackList(userID, trackID)
	if err != nil {
		return nil, err
	}

	userTrackList, err := s.repo.GetUserTrackList(userID)
	if err != nil {
		return nil, err
	}

	return userTrackList, nil
}

func (s Service) GetAllPlaylists() ([]model.Playlist, [][]model.Track, error) {
	playlists, playlistsTracks, err := s.repo.GetAllPlaylists()
	if err != nil {
		return nil, nil, err
	}

	return playlists, playlistsTracks, nil
}

func (s Service) GetUserPlaylists(userID string) ([]model.Playlist, [][]model.Track, error) {
	playlists, playlistsTracks, err := s.repo.GetUserPlaylists(userID)
	if err != nil {
		return nil, nil, err
	}

	return playlists, playlistsTracks, nil
}

func (s Service) CreateNewPlaylist(title, createdByID string, trackList []string) (model.Playlist, []model.Track, error) {
	err := model.ValidatePlaylist(title)
	if err != nil {
		return model.Playlist{}, nil, err
	}

	id, err := s.repo.CreateNewPlaylist(title, createdByID, trackList)
	if err != nil {
		return model.Playlist{}, nil, err
	}

	playlist, playlistTracks, err := s.repo.GetPlaylistByID(id)
	if err != nil {
		return model.Playlist{}, nil, err
	}

	return playlist, playlistTracks, nil
}

func (s Service) GetPlaylistByID(playlistID string) (model.Playlist, []model.Track, error) {
	playlist, playlistTracks, err := s.repo.GetPlaylistByID(playlistID)
	if err != nil {
		return model.Playlist{}, nil, err
	}

	return playlist, playlistTracks, nil
}

func (s Service) AddTracksToPlaylist(userID, playlistID string, trackList []string) (model.Playlist, []model.Track, error) {
	err := s.repo.AddTracksToPlaylist(userID, playlistID, trackList)
	if err != nil {
		return model.Playlist{}, nil, err
	}

	playlist, playlistTracks, err := s.repo.GetPlaylistByID(playlistID)
	if err != nil {
		return model.Playlist{}, nil, err
	}

	return playlist, playlistTracks, nil
}

func (s Service) RemoveTracksFromPlaylist(userID, playlistID string, trackList []string) (model.Playlist, []model.Track, error) {
	err := s.repo.RemoveTracksFromPlaylist(userID, playlistID, trackList)
	if err != nil {
		return model.Playlist{}, nil, err
	}

	playlist, playlistTracks, err := s.repo.GetPlaylistByID(playlistID)
	if err != nil {
		return model.Playlist{}, nil, err
	}

	return playlist, playlistTracks, nil
}

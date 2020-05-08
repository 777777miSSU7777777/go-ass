package api

import (
	"github.com/777777miSSU7777777/go-ass/model"
)

type AudioResponse struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

type AddAudioResponse AudioResponse

type GetAudioListResponse []model.Audio

type GetAudioByIDResponse AudioResponse

type UpdateAudioByIDResponse AudioResponse

type DeleteAudioByIDResponse struct {
}

type SignUpResponse struct{}

type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse SignInResponse

type SignOutResponse struct{}

type AddAudioToUserAudioListResponse struct{}

type DeleteAudioFromUserAudioListResponse struct{}

type GetUserAudioListResponse GetAudioListResponse

type AudioPlaylistResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	TrackList []AudioResponse `json:"tracklist"`
}

type GetAllAudioPlayListsResponse []AudioPlaylistResponse

type CreateNewPlaylistResponse AudioPlaylistResponse

type GetAudioPlaylistByIDResponse AudioPlaylistResponse

type AddAudioListToPlaylistResponse AudioPlaylistResponse

type DeleteAudioListFromPlaylistResponse AudioPlaylistResponse
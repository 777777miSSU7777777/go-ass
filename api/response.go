package api

import (
	"github.com/777777miSSU7777777/go-ass/model"
)

type AddAudioResponse struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

type GetAudioListResponse []model.Audio

type GetAudioByIDResponse AddAudioResponse

type UpdateAudioByIDResponse AddAudioResponse

type DeleteAudioByIDResponse struct {
}

type SignUpResponse struct{}

type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse SignInResponse

type SignOutResponse struct{}

type AddAudioToUserAudioListResponse struct {}

type DeleteAudioFromUserAudioListResponse struct {}

type GetUserAudioListResponse GetAudioListResponse
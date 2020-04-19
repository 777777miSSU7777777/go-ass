package api

import (
	"github.com/777777miSSU7777777/go-ass/model"
)

type AddAudioResponse struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

type GetAudioListResponse struct {
	Audio []model.Audio `json:"audio"`
}

type GetAudioByIDResponse AddAudioResponse

type UpdateAudioByIDResponse AddAudioResponse

type DeleteAudioByIDResponse struct {
}

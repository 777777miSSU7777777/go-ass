package api

import (
	"github.com/777777miSSU7777777/go-ass/model"
)

type AddAudioResponse struct {
	ID     int64  `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

type GetAllAudioResponse struct {
	Audio []model.Audio `json:"audio"`
}

type GetAudioByIDResponse AddAudioResponse

type GetAudioByKeyResponse GetAllAudioResponse

type UpdateAudioByIDResponse AddAudioResponse

type DeleteAudioByIDResponse struct {
}

package model

import (
	"errors"
)

var AudioAuthorEmpty = errors.New("audio author name is empty")
var AudioTitleEmpty = errors.New("audio title is empty")

type Audio struct {
	ID     int64  `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

func ValidateAudio(author, title string) error {
	if author == "" {
		return AudioAuthorEmpty
	}
	if title == "" {
		return AudioTitleEmpty
	}

	return nil
}

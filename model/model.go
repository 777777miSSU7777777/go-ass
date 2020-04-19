package model

import (
	"errors"
)

var AudioAuthorEmpty = errors.New("audio author name is empty")
var AudioTitleEmpty = errors.New("audio title is empty")

type Audio struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Author string `json:"author" bson:"author"`
	Title  string `json:"title" bson:"title"`
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

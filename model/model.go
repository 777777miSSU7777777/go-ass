package model

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var TrackAuthorEmpty = errors.New("track author name is empty")
var TrackTitleEmpty = errors.New("track title is empty")
var UserEmailEmpty = errors.New("user email is empty")
var UserEmailLength = errors.New("user email must be between 20 - 40 symbols length")
var UserNameEmpty = errors.New("user name is empty")
var UserNameLength = errors.New("user name length must be between 5 - 16 symbols length")
var UserPasswordEmpty = errors.New("user password is empty")
var UserPasswordLength = errors.New("user password must be between 8 - 40 symbols length")
var PlaylistTitleEmpty = errors.New("playlist title is empty")

type Track struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author       string             `json:"author" bson:"author"`
	Title        string             `json:"title" bson:"title"`
	UploadedByID primitive.ObjectID `json:"uploadedByID" bson:"uploadedByID"`
}

func ValidateTrack(author, title string) error {
	if author == "" {
		return TrackAuthorEmpty
	}
	if title == "" {
		return TrackTitleEmpty
	}

	return nil
}

type User struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty"`
	Email         string               `bson:"email"`
	Name          string               `bson:"name"`
	Password      string               `bson:"password"`
	RefreshTokens []string             `bson:"refresh_tokens"`
	TrackList     []primitive.ObjectID `bson:"tracklist"`
	Playlists     []primitive.ObjectID `bson:"playlists"`
}

func ValidateUser(email, name, password string) error {
	if email == "" {
		return UserEmailEmpty
	}
	if len(email) < 20 && len(email) > 40 {
		return UserEmailLength
	}

	if name == "" {
		return UserNameEmpty
	}
	if len(name) < 5 && len(name) > 16 {
		return UserNameLength
	}

	if password == "" {
		return UserPasswordEmpty
	}
	if len(password) < 8 && len(password) > 40 {
		return UserPasswordLength
	}

	return nil
}

type Playlist struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Title       string               `bson:"title"`
	TrackList   []primitive.ObjectID `bson:"tracklist"`
	CreatedByID primitive.ObjectID   `bson:"createdByID"`
}

func ValidatePlaylist(title string) error {
	if title == "" {
		return PlaylistTitleEmpty
	}

	return nil
}

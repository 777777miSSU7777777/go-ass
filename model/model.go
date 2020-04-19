package model

import (
	"errors"
)

var AudioAuthorEmpty = errors.New("audio author name is empty")
var AudioTitleEmpty = errors.New("audio title is empty")
var UserEmailEmpty = errors.New("user email is empty")
var UserEmailLength = errors.New("user email must be between 20 - 40 symbols length")
var UserNameEmpty = errors.New("user name is empty")
var UserNameLength = errors.New("user name length must be between 5 - 16 symbols length")
var UserPasswordEmpty = errors.New("user password is empty")
var UserPasswordLength = errors.New("user password must be between 8 - 40 symbols length")

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

type User struct {
	ID string `json:"id" bson:"_id,omitempty"`
	Email string `json:"email" bson:"email"`
	Name string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
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
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

func (s Service) AddAudio(author, title, uploadedByID string) (model.Audio, error) {
	err := model.ValidateAudio(author, title)
	if err != nil {
		return model.Audio{}, err
	}

	id, err := s.repo.AddAudio(author, title, uploadedByID)
	if err != nil {
		return model.Audio{}, err
	}

	audio, err := s.repo.GetAudioByID(id)
	if err != nil {
		return model.Audio{}, err
	}

	return audio, nil
}

func (s Service) GetAllAudio() ([]model.Audio, error) {
	audio, err := s.repo.GetAllAudio()
	if err != nil {
		return nil, err
	}

	return audio, nil
}

func (s Service) GetAudioByID(id string) (model.Audio, error) {
	audio, err := s.repo.GetAudioByID(id)
	if err != nil {
		return model.Audio{}, err
	}

	return audio, nil
}

func (s Service) GetAudioByKey(key string) ([]model.Audio, error) {
	audio, err := s.repo.GetAudioByKey(key)
	if err != nil {
		return nil, err
	}

	return audio, nil
}

func (s Service) UpdateAudioByID(id string, author, title string) (model.Audio, error) {
	err := model.ValidateAudio(author, title)
	if err != nil {
		return model.Audio{}, err
	}

	err = s.repo.UpdateAudioByID(id, author, title)
	if err != nil {
		return model.Audio{}, err
	}

	audio, err := s.repo.GetAudioByID(id)
	if err != nil {
		return model.Audio{}, err
	}

	return audio, nil
}

func (s Service) DeleteAudioByID(id string) error {
	err := s.repo.DeleteAudioByID(id)
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
			user.ID,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Second * time.Duration(1800)).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
		accessToken, err := token.SignedString([]byte(repository.SecretKey))
		if err != nil {
			return "", "", fmt.Errorf("error while signing user refresh token: %v", err)
		}

		refreshToken, err := s.repo.AddRefreshToken(user.ID)
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

package service

import (
	"github.com/777777miSSU7777777/go-ass/model"
	"github.com/777777miSSU7777777/go-ass/repository"
)

type Service struct {
	repo repository.Repository
}

func New(r repository.Repository) Service {
	return Service{r}
}

func (s Service) AddAudio(author, title string) (model.Audio, error) {
	err := model.ValidateAudio(author, title)
	if err != nil {
		return model.Audio{}, err
	}

	id, err := s.repo.AddAudio(author, title)
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
		return  err
	}

	return nil
}

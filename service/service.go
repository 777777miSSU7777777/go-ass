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

func (s Service) GetLastAudioID() (int64, error) {
	id, err := s.repo.GetLastID("audio")

	if err != nil {
		return -1, err
	}

	return id, nil
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

func (s Service) GetAudioByID(id int64) (model.Audio, error) {
	audio, err := s.repo.GetAudioByID(id)
	if err != nil {
		return model.Audio{}, err
	}

	return audio, nil
}

func (s Service) GetAudioByAuthor(author string) ([]model.Audio, error) {
	audio, err := s.repo.GetAudioByAuthor(author)
	if err != nil {
		return nil, err
	}

	return audio, nil
}

func (s Service) GetAudioByTitle(title string) ([]model.Audio, error) {
	audio, err := s.repo.GetAudioByTitle(title)
	if err != nil {
		return nil, err
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

func (s Service) UpdateAudioByID(id int64, author, title string) (model.Audio, error) {
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

func (s Service) DeleteAudioByID(id int64) error {
	err := s.repo.DeleteAudioByID(id)
	if err != nil {
		return err
	}

	return nil
}

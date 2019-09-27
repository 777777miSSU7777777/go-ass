package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/777777miSSU7777777/go-ass/model"
	_ "github.com/go-sql-driver/mysql"
)

var AudioNotFoundError = errors.New("audio not found error")

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return Repository{db}
}

func (r Repository) GetLastAudioID() (int64, error) {
	row := r.db.QueryRow("SELECT id FROM audio GROUP BY id DESC LIMIT 1")
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("get last audio id error: %v", err)
	}

	return id, nil
}

func (r Repository) AddAudio(author, title string) (int64, error) {
	result, err := r.db.Exec("INSERT INTO audio(author, title) VALUES(?,?)", author, title)
	if err != nil {
		return -1, fmt.Errorf("add audio error: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("add audio error: %v", err)
	}

	return id, nil
}

func (r Repository) GetAllAudio() ([]model.Audio, error) {
	rows, err := r.db.Query("SELECT * FROM audio")
	if err != nil {
		return nil, fmt.Errorf("get all audio error: %v", err)
	}
	defer rows.Close()

	if err == sql.ErrNoRows {
		return nil, AudioNotFoundError
	}

	audio := []model.Audio{}
	for rows.Next() {
		track := model.Audio{}
		err := rows.Scan(&track.ID, &track.Author, &track.Title)
		if err != nil {
			return nil, fmt.Errorf("audio scan error: %v", err)
		}
		audio = append(audio, track)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("scan audio error: %v", err)
	}

	return audio, nil
}

func (r Repository) GetAudioByID(id int64) (model.Audio, error) {
	row := r.db.QueryRow("SELECT * FROM audio WHERE id=?", id)
	audio := model.Audio{}
	err := row.Scan(&audio.ID, &audio.Author, &audio.Title)
	if err != nil {
		if err != sql.ErrNoRows {
			return model.Audio{}, AudioNotFoundError
		}
		return model.Audio{}, fmt.Errorf("get audio by id error: %v", err)
	}

	return audio, nil
}

func (r Repository) GetAudioByAuthor(author string) ([]model.Audio, error) {
	rows, err := r.db.Query("SELECT * FROM audio WHERE author LIKE '%?%'", author)
	if err != nil {
		return nil, fmt.Errorf("get audio by author error: %v", err)
	}
	defer rows.Close()

	if err == sql.ErrNoRows {
		return nil, AudioNotFoundError
	}

	audio := []model.Audio{}
	for rows.Next() {
		track := model.Audio{}
		err := rows.Scan(&track.ID, &track.Author, &track.Title)
		if err != nil {
			return nil, fmt.Errorf("audio scan error: %v", err)
		}
		audio = append(audio, track)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("audio scan error: %v", err)
	}

	return audio, nil
}

func (r Repository) GetAudioByTitle(title string) ([]model.Audio, error) {
	rows, err := r.db.Query("SELECT * FROM audio WHERE title LIKE '%?%'", title)
	if err != nil {
		return nil, fmt.Errorf("get audio by title error: %v", err)
	}
	defer rows.Close()

	if err == sql.ErrNoRows {
		return nil, AudioNotFoundError
	}

	audio := []model.Audio{}
	for rows.Next() {
		track := model.Audio{}
		err := rows.Scan(&track.ID, &track.Author, &track.Title)
		if err != nil {
			return nil, fmt.Errorf("audio scan error: %v", err)
		}
		audio = append(audio, track)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("audio scan error: %v", err)
	}

	return audio, nil
}

func (r Repository) UpdateAudioByID(id int64, author, title string) error {
	result, err := r.db.Exec("UPDATE audio SET author=?, title=? WHERE id=?", author, title, id)
	if err != nil {
		return fmt.Errorf("update audio by id error: %v", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("update audio by id error: %v", err)
	}
	if affected == 0 {
		return AudioNotFoundError
	}

	return nil
}

func (r Repository) DeleteAudioByID(id int64) error {
	result, err := r.db.Exec("DELETE FROM audio WHERE id=?", id)
	if err != nil {
		return fmt.Errorf("delete audio by id error: %v", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("delete audio by id error: %v", err)
	}
	if affected == 0 {
		return AudioNotFoundError
	}

	return nil
}

// func (r Repository) GetAllAuthors() []model.Author {

// }

// func (r Repository) GetAuthorByID(id int64) model.Author {

// }

// func (r Repository) GetAuthorByName(name string) model.Author {

// }

package model

import (
	"github.com/jinzhu/gorm"
)

type GenreTracks struct {
	gorm.Model
	GenreID int64 `gorm: "column: genre_id; type: bigint; primary_key; auto_increment: false; unique; not null"`
	TrackID int64 `gorm: "column: track_id; type: bigint; primary_key; auto_increment: false; unique; not null"`
}

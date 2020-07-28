package model

import (
	"github.com/jinzhu/gorm"
)

type Track struct {
	gorm.Model
	TrackID      int64  `gorm: "column: track_id; type: bigint; primary_key; unique; not null"`
	TrackTitle   string `gorm: "column: track_title; type: nvarchar(50); not null"`
	ArtistID     int64  `gorm: "column: artist_id; type: bigint; not null"`
	GenreID      int64  `gorm: "column: genre_id; type: bigint;"`
	UploadedByID int64  `gorm: "column: uploaded_by_id: type: bigint; not null"`
}

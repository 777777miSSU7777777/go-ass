package model

import (
	"github.com/jinzhu/gorm"
)

type Artist struct {
	gorm.Model
	ArtistID   int64  `gorm: "column: artist_id; type: bigint; primary_key; unique; not null"`
	ArtistName string `gorm: "column: artist_name; type: nvarchar(50); unique; not null"`
}

package model

import (
	"github.com/jinzhu/gorm"
)

type Playlist struct {
	gorm.Model
	PlaylistID    int64  `gorm: "column: playlist_id; type: bigint; primary_key; unique; not null"`
	PlaylistTitle string `gorm: "column: playlist_title; type: nvarchar(50); unique; not null"`
	CreatedByID   int64  `gorm: "column: created_by_id; type: bigint; not null"`
}

package model

import (
	"github.com/jinzhu/gorm"
)

type PlaylistTracks struct {
	gorm.Model
	PlaylistID int64 `gorm: "column: playlist_id; type: bigint; primary_key; auto_increment: false; unique; not null"`
	TrackID    int64 `gorm: "column: track_id; type: bigint; primary_key; auto_increment: false; unique; not null"`
}

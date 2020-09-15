package model

import (
	"github.com/jinzhu/gorm"
)

type UserPlaylists struct {
	gorm.Model
	UserID     int64 `gorm: "column: user_id; type: bigint; primary_key; auto_increment: false; unique; not null"`
	PlaylistID int64 `gorm: "column: playlist_id type: bigint; primary_key; auto_increment: false; unique; not null"`
}

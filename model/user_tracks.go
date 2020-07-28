package model

import (
	"github.com/jinzhu/gorm"
)

type UserTracks struct {
	gorm.Model
	UserID  int64 `gorm: "type: bigint; primary_key; auto_increment: false; unique; not null"`
	TrackID int64 `gorm: "type: bigint; primary_key; auto_increment: false; unique; not null"`
}

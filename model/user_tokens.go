package model

import (
	"github.com/jinzhu/gorm"
)

type UserTokens struct {
	gorm.Model
	UserID int64  `gorm: "type: bigint; primary_key; not null"`
	Token  string `gorm: "type: tinytext; primary_key; not null"`
}

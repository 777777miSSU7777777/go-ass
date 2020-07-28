package model

import (
	"github.com/jinzhu/gorm"
)

type Genre struct {
	gorm.Model
	GenreID    int64  `gorm: "column: genre_id; type: bigint; primary_key; unique; not null"`
	GenreTitle string `gorm: "column: genre_title; type: nvarchar(50); unique; not null"`
}

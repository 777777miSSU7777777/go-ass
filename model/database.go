package model

import (
	"github.com/jinzhu/gorm"
)

type Artist struct {
	gorm.Model
	ArtistID   int64  `gorm: "column: artist_id; type: bigint; primary_key; unique; not null"`
	ArtistName string `gorm: "column: artist_name; type: nvarchar(50); unique; not null"`
}

type GenreTracks struct {
	gorm.Model
	GenreID int64 `gorm: "column: genre_id; type: bigint; primary_key; auto_increment: false; unique; not null"`
	TrackID int64 `gorm: "column: track_id; type: bigint; primary_key; auto_increment: false; unique; not null"`
}

type Genre struct {
	gorm.Model
	GenreID    int64  `gorm: "column: genre_id; type: bigint; primary_key; unique; not null"`
	GenreTitle string `gorm: "column: genre_title; type: nvarchar(50); unique; not null"`
}

type PlaylistTracks struct {
	gorm.Model
	PlaylistID int64 `gorm: "column: playlist_id; type: bigint; primary_key; auto_increment: false; unique; not null"`
	TrackID    int64 `gorm: "column: track_id; type: bigint; primary_key; auto_increment: false; unique; not null"`
}

type Playlist struct {
	gorm.Model
	PlaylistID    int64  `gorm: "column: playlist_id; type: bigint; primary_key; unique; not null"`
	PlaylistTitle string `gorm: "column: playlist_title; type: nvarchar(50); unique; not null"`
	CreatedByID   int64  `gorm: "column: created_by_id; type: bigint; not null"`
}

type Track struct {
	gorm.Model
	TrackID      int64  `gorm: "column: track_id; type: bigint; primary_key; unique; not null"`
	TrackTitle   string `gorm: "column: track_title; type: nvarchar(50); not null"`
	ArtistID     int64  `gorm: "column: artist_id; type: bigint; not null"`
	GenreID      int64  `gorm: "column: genre_id; type: bigint;"`
	UploadedByID int64  `gorm: "column: uploaded_by_id: type: bigint; not null"`
}

type UserPlaylists struct {
	gorm.Model
	UserID     int64 `gorm: "column: user_id; type: bigint; primary_key; auto_increment: false; unique; not null"`
	PlaylistID int64 `gorm: "column: playlist_id type: bigint; primary_key; auto_increment: false; unique; not null"`
}

type UserTokens struct {
	gorm.Model
	UserID int64  `gorm: "type: bigint; primary_key; not null"`
	Token  string `gorm: "type: tinytext; primary_key; not null"`
}

type UserTracks struct {
	gorm.Model
	UserID  int64 `gorm: "type: bigint; primary_key; auto_increment: false; unique; not null"`
	TrackID int64 `gorm: "type: bigint; primary_key; auto_increment: false; unique; not null"`
}

type User struct {
	gorm.Model
	UserID   int64  `gorm: "column: user_id; type: bigint; primary_key; unique; not null"`
	Role     string `gorm: "column: role; type: nvarchar(50); not null;"`
	Email    string `gorm: "column: email; type: nvarchar(50); unique; not null"`
	Username string `gorm: "column: username; type: nvarchar(50); unique; not null"`
	Password string `gorm: "column: password; type: tinytext; not null"`
}

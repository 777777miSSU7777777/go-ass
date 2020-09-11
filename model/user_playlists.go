package model

type UserPlaylists struct {
	gorm.model
	UserID     int64 `gorm: "column: user_id; type: bigint; primary_key; auto_increment: false; unique; not null"`
	PlaylistID int64 `gorm: "column: playlist_id type: bigint; primary_key; auto_increment: false; unique; not null"`
}

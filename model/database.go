package model

type Artist struct {
	ArtistID   string `gorm:"column:artist_id;type:uuid;primaryKey;unique;not null"`
	ArtistName string `gorm:"column:artist_name;type:varchar(50);unique;not null"`
}

type GenreTracks struct {
	GenreID string `gorm:"column:genre_id;type:uuid;primaryKey;not null"`
	TrackID string `gorm:"column:track_id;type:uuid;primaryKey;not null"`
}

type Genre struct {
	GenreID    string `gorm:"column:genre_id;type:uuid;primaryKey;unique;not null"`
	GenreTitle string `gorm:"column:genre_title;type:varchar(50);unique;not null"`
}

type PlaylistTracks struct {
	PlaylistID string `gorm:"column:playlist_id;type:uuid;primaryKey;not null"`
	TrackID    string `gorm:"column:track_id;type:uuid;primaryKey;not null"`
}

type Playlist struct {
	PlaylistID    string `gorm:"column:playlist_id;type:uuid;primaryKey;unique;not null"`
	PlaylistTitle string `gorm:"column:playlist_title;type:varchar(50);not null"`
	CreatedByID   string `gorm:"column:created_by_id;type:uuid;not null"`
}

type Track struct {
	TrackID      string `gorm:"column:track_id;type:uuid;primaryKey;unique;not null"`
	TrackTitle   string `gorm:"column:track_title;type:varchar(50);not null"`
	ArtistID     string `gorm:"column:artist_id;type:uuid;not null"`
	GenreID      string `gorm:"column:genre_id;type:uuid;not null"`
	UploadedByID string `gorm:"column:uploaded_by_id;type:uuid;not null"`
}

type UserPlaylists struct {
	UserID     string `gorm:"column:user_id;type:uuid;primaryKey;not null"`
	PlaylistID string `gorm:"column:playlist_id;type:uuid;primaryKey;not null"`
}

type UserTokens struct {
	UserID string `gorm:"column:user_id;type:uuid;primaryKey;not null"`
	Token  string `gorm:"coulmn:token;type:text;primaryKey;not null"`
}

type UserTracks struct {
	UserID  string `gorm:"column:user_id;type:uuid;primaryKey;not null"`
	TrackID string `gorm:"coulmn:track_id;type:uuid;primaryKey;not null"`
}

type User struct {
	UserID   string `gorm:"column:user_id;type:uuid;primaryKey;unique;not null"`
	Role     string `gorm:"column:role;type:varchar(50);not null;"`
	Email    string `gorm:"column:email;type:varchar(50);unique;not null"`
	Username string `gorm:"column:username;type:varchar(50);unique;not null"`
	Password string `gorm:"column:password;type:text;not null"`
}

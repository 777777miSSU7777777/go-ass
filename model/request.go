package model

type UpdateTrackByIDRequest struct {
	TrackTitle string `json:"trackTitle"`
	ArtistID   string `json:"artistId`
	GenreID    string `json:"genreId`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken`
}

type SignOutRequest RefreshTokenRequest

type AddTracksToUserListRequest struct {
	TrackList []string `json:"trackList"`
}

type DeleteTracksFromUserListRequest AddTracksToUserListRequest

type CreateNewPlaylistRequest struct {
	PlaylistTitle string   `json:"playlistTitle`
	TrackList     []string `json:"trackList`
}

type AddTracksToPlaylistRequest AddTracksToUserListRequest

type DeleteTracksFromPlaylistRequest AddTracksToPlaylistRequest

type AddPlaylistsToUserListRequest struct {
	Playlists []string `json:"playlists`
}

type DeletePlaylistsFromUserListRequest AddPlaylistsToUserListRequest

package model

type UpdateTrackByIDRequest struct {
	TrackTitle string `json: "trackTitle"`
	ArtistID   int64  `json: "artistId`
	GenreID    int64  `json: "genreId`
}

type SignUpRequest struct {
	Email    string `json: "email"`
	Username string `json: "username"`
	Password string `json: "password"`
}

type SignInRequest struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json: "refreshToken`
}

type SignOutRequest RefreshTokenRequest

type AddTracksToUserListRequest struct {
	TrackList []int64 `json: "trackList"`
}

type DeleteTracksFromUserListRequest AddTracksToUserListRequest

type CreateNewPlaylistRequest struct {
	PlaylistTitle string  `json: "playlistTitle`
	TrackList     []int64 `json: "trackList`
}

type AddTracksToPlaylistRequest AddTracksToUserListRequest

type DeleteTracksFromPlaylistRequest AddTracksToPlaylistRequest

type AddPlaylistsToUserListRequest struct {
	Playlists []int64 `json: "playlists`
}

type DeletePlaylistsFromUserListRequest AddPlaylistsToUserListRequest

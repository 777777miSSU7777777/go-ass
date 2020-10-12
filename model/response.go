package model

type ArtistResponse struct {
	ID   string `json:"artistId"`
	Name string `json:"artistName"`
}

type TrackResponse struct {
	ID     string `json:"trackId"`
	Title  string `json:"trackTitle"`
	Artist string `json:"artistName"`
}

type PlaylistResponse struct {
	ID        string          `json:"playlistId"`
	Title     string          `json:"playlistTitle"`
	TrackList []TrackResponse `json:"trackList"`
}

type SignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse SignInResponse

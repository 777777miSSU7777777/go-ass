package model

type ArtistResponse struct {
	ID   string `json: "artistId"`
	Name string `json: "artistName"`
}

type TrackResponse struct {
	ID     int64  `json: "trackID"`
	Title  string `json: "trackTitle"`
	Artist string `json: "artistName"`
}

type PlaylistResponse struct {
	ID        int64           `json: "playlistId"`
	Title     string          `json: "playlistTitle"`
	TrackList []TrackResponse `json: "trackList"`
}

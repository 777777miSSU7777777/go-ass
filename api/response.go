package api

type TrackResponse struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

type AddTrackResponse TrackResponse

type GetAllTracksResponse []TrackResponse

type GetTrackByIDResponse TrackResponse

type UpdateTrackByIDResponse TrackResponse

type DeleteTrackByIDResponse struct{}

type SignUpResponse struct{}

type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse SignInResponse

type SignOutResponse struct{}

type AddTrackToUserTrackListResponse GetAllTracksResponse

type RemoveTrackFromUserTrackListResponse GetAllTracksResponse

type GetUserTrackListResponse GetAllTracksResponse

type PlaylistResponse struct {
	ID        string          `json:"id"`
	Title     string          `json:"title"`
	TrackList []TrackResponse `json:"tracklist"`
}

type GetAllPlaylistsResponse []PlaylistResponse

type GetUserPlaylistsResponse []PlaylistResponse

type CreateNewPlaylistResponse PlaylistResponse

type GetPlaylistByIDResponse PlaylistResponse

type AddTracksToPlaylistResponse PlaylistResponse

type RemoveTracksFromPlaylistResponse PlaylistResponse

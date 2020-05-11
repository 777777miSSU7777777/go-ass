package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/777777miSSU7777777/go-ass/model"
	"github.com/777777miSSU7777777/go-ass/repository"
	"github.com/777777miSSU7777777/go-ass/service"
)

type ErrorResponse struct {
	Type  string `json:"type"`
	Error string `json:"error"`
}

var BodyParseError = "BODY PARSE ERROR"
var IDParseError = "ID PARSE ERROR"
var ValidationError = "VALIDATION ERROR"
var NotFoundError = "NOT FOUND ERROR"
var ServiceError = "SERVICE ERROR"
var QueryStringError = "QUERY STRING ERROR"
var InternalServerError = "INTERNAL SERVER ERROR"

func writeError(w http.ResponseWriter, statusCode int, errType string, err error) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Type: errType, Error: err.Error()})
}

type API struct {
	svc service.Service
	m   FileManager
}

func NewApi(svc service.Service, m FileManager) API {
	return API{svc, m}
}

func (a API) AddTrack(w http.ResponseWriter, r *http.Request) {
	author, title := r.FormValue("author"), r.FormValue("title")

	userID := r.Context().Value("userID").(string)

	newTrack, err := a.svc.AddTrack(author, title, userID)
	if err != nil {
		if err.Error() == model.TrackAuthorEmpty.Error() || err.Error() == model.TrackTitleEmpty.Error() {
			writeError(w, 400, ValidationError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		_ = a.m.Delete(w, newTrack.ID.Hex())
		return
	} else {
		err = a.m.Upload(w, r, newTrack.ID.Hex())
		if err != nil {
			_ = a.svc.DeleteTrackByID(newTrack.ID.Hex())
			_ = a.m.Delete(w, newTrack.ID.Hex())
			return
		}
	}

	_ = json.NewEncoder(w).Encode(AddTrackResponse{newTrack.ID.Hex(), newTrack.Author, newTrack.Title})
}

func (a API) GetAllTracks(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	var tracks []model.Track
	var err error
	if key != "" {
		tracks, err = a.svc.GetTracksByKey(key)
		if err != nil {
			if err.Error() == repository.TrackNotFoundError.Error() {
				writeError(w, 400, NotFoundError, err)
			} else {
				writeError(w, 400, ServiceError, err)
			}
			return
		}

	} else {
		tracks, err = a.svc.GetAllTracks()
		if err != nil {
			if err.Error() == repository.TrackNotFoundError.Error() {
				writeError(w, 404, NotFoundError, err)
			} else {
				writeError(w, 400, ServiceError, err)
			}
			return
		}
	}

	resp := GetAllTracksResponse{}
	for _, track := range tracks {
		resp = append(resp, TrackResponse{ID: track.ID.Hex(), Author: track.Author, Title: track.Title})
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (a API) GetTrackByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackID := vars["id"]

	track, err := a.svc.GetTrackByID(trackID)
	if err != nil {
		if err.Error() == repository.TrackNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(GetTrackByIDResponse{track.ID.Hex(), track.Author, track.Title})
}

func (a API) UpdateTrackByID(w http.ResponseWriter, r *http.Request) {
	var req UpdateTrackByIDRequest
	vars := mux.Vars(r)
	trackID := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	updatedTrack, err := a.svc.UpdateTrackByID(trackID, req.Author, req.Title)
	if err != nil {
		if err.Error() == repository.TrackNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(UpdateTrackByIDResponse{updatedTrack.ID.Hex(), updatedTrack.Author, updatedTrack.Title})
}

func (a API) DeleteTrackByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackID := vars["id"]

	err := a.m.Delete(w, trackID)
	if err != nil {
		return
	}

	err = a.svc.DeleteTrackByID(trackID)
	if err != nil {
		if err.Error() == repository.TrackNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(DeleteTrackByIDResponse{})
}

func (a API) SignUp(w http.ResponseWriter, r *http.Request) {
	var req SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	err = a.svc.SignUp(req.Email, req.Name, req.Password)
	if err != nil {
		writeError(w, 400, ServiceError, err)
		return
	}

	_ = json.NewEncoder(w).Encode(SignUpResponse{})
}

func (a API) SignIn(w http.ResponseWriter, r *http.Request) {
	var req SignInRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	accessToken, refreshToken, err := a.svc.SignIn(req.Email, req.Password)
	if err != nil {
		if err.Error() == repository.UserNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(SignInResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

func (a API) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshTokenRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	accessToken, refreshToken, err := a.svc.RefreshToken(req.RefreshToken)
	if err != nil {
		if err.Error() == repository.RefreshTokenNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(SignInResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

func (a API) SignOut(w http.ResponseWriter, r *http.Request) {
	var req SignOutRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	err = a.svc.SignOut(req.RefreshToken)
	if err != nil {
		if err.Error() == repository.RefreshTokenNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(SignOutResponse{})
}

func (a API) GetUserTrackList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	userTrackList, err := a.svc.GetUserTrackList(userID)
	if err != nil {
		if err.Error() == repository.UserNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	resp := GetUserTrackListResponse{}
	for _, track := range userTrackList {
		resp = append(resp, TrackResponse{ID: track.ID.Hex(), Author: track.Author, Title: track.Title})
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (a API) AddTrackToUserTrackList(w http.ResponseWriter, r *http.Request) {
	var req AddTrackToUserTrackListRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	userID := r.Context().Value("userID").(string)

	resp, err := a.svc.AddTrackToUserTrackList(userID, req.TrackID)
	if err != nil {
		writeError(w, 400, ServiceError, err)
		return
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (a API) RemoveTrackFromUserTrackList(w http.ResponseWriter, r *http.Request) {
	var req RemoveTrackFromUserTrackListRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	userID := r.Context().Value("userID").(string)

	resp, err := a.svc.RemoveTrackFromUserTrackList(userID, req.TrackID)
	if err != nil {
		writeError(w, 400, ServiceError, err)
		return
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (a API) GetAllPlaylists(w http.ResponseWriter, r *http.Request) {
	playlists, playlistsTracks, err := a.svc.GetAllPlaylists()
	if err != nil {
		writeError(w, 400, ServiceError, err)
		return
	}

	resp := GetAllPlaylistsResponse{}
	for playlistIndex, playlist := range playlists {
		trackList := []TrackResponse{}
		for _, playlistTrack := range playlistsTracks[playlistIndex] {
			trackList = append(trackList, TrackResponse{ID: playlistTrack.ID.Hex(), Author: playlistTrack.Author, Title: playlistTrack.Title})
		}
		resp = append(resp, PlaylistResponse{ID: playlist.ID.Hex(), Title: playlist.Title, TrackList: trackList})
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (a API) GetUserPlaylists(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	playlists, playlistsTracks, err := a.svc.GetUserPlaylists(userID)
	if err != nil {
		writeError(w, 400, ServiceError, err)
		return
	}

	resp := GetUserPlaylistsResponse{}
	for playlistIndex, playlist := range playlists {
		trackList := []TrackResponse{}
		for _, playlistTrack := range playlistsTracks[playlistIndex] {
			trackList = append(trackList, TrackResponse{ID: playlistTrack.ID.Hex(), Author: playlistTrack.Author, Title: playlistTrack.Title})
		}
		resp = append(resp, PlaylistResponse{ID: playlist.ID.Hex(), Title: playlist.Title, TrackList: trackList})
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (a API) CreateNewPlaylist(w http.ResponseWriter, r *http.Request) {
	var req CreateNewPlaylistRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	userID := r.Context().Value("userID").(string)

	playlist, playlistTracks, err := a.svc.CreateNewPlaylist(req.Title, userID, req.TrackList)
	if err != nil {
		writeError(w, 400, ServiceError, err)
		return
	}

	resp := PlaylistResponse{ID: playlist.ID.Hex(), Title: playlist.Title}
	for _, playlistTrack := range playlistTracks {
		track := TrackResponse{ID: playlistTrack.ID.Hex(), Author: playlistTrack.Author, Title: playlistTrack.Title}
		resp.TrackList = append(resp.TrackList, track)
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (a API) GetPlaylistByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistID := vars["id"]

	playlist, playlistTracks, err := a.svc.GetPlaylistByID(playlistID)
	if err != nil {
		writeError(w, 400, ServiceError, err)
		return
	}

	resp := GetPlaylistByIDResponse{ID: playlist.ID.Hex(), Title: playlist.Title}
	for _, playlistTrack := range playlistTracks {
		track := TrackResponse{ID: playlistTrack.ID.Hex(), Author: playlistTrack.Author, Title: playlistTrack.Title}
		resp.TrackList = append(resp.TrackList, track)
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (a API) DeletePlaylistByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistID := vars["id"]

	userID := r.Context().Value("userID").(string)

	err := a.svc.DeletePlaylistByID(playlistID, userID)
	if err != nil {
		if err.Error() == repository.PlaylistNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
			return
		} else {
			writeError(w, 400, ServiceError, err)
			return
		}
	}

	_ = json.NewEncoder(w).Encode(DeletePlaylistByIDResponse{})
}

func (a API) AddTracksToPlaylist(w http.ResponseWriter, r *http.Request) {
	var req AddTracksToPlaylistRequest

	vars := mux.Vars(r)
	playlistID := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	userID := r.Context().Value("userID").(string)

	playlist, playlistTracks, err := a.svc.AddTracksToPlaylist(userID, playlistID, req.TrackList)
	if err != nil {
		writeError(w, 400, ServiceError, err)
		return
	}

	resp := AddTracksToPlaylistResponse{ID: playlist.ID.Hex(), Title: playlist.Title}
	for _, playlistTrack := range playlistTracks {
		track := TrackResponse{ID: playlistTrack.ID.Hex(), Author: playlistTrack.Author, Title: playlistTrack.Title}
		resp.TrackList = append(resp.TrackList, track)
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (a API) RemoveTracksFromPlaylist(w http.ResponseWriter, r *http.Request) {
	var req RemoveTracksFromPlaylistRequest

	vars := mux.Vars(r)
	playlistID := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	userID := r.Context().Value("userID").(string)

	playlist, playlistTracks, err := a.svc.RemoveTracksFromPlaylist(userID, playlistID, req.TrackList)
	if err != nil {
		writeError(w, 400, ServiceError, err)
		return
	}

	resp := RemoveTracksFromPlaylistResponse{ID: playlist.ID.Hex(), Title: playlist.Title}
	for _, playlistTrack := range playlistTracks {
		track := TrackResponse{ID: playlistTrack.ID.Hex(), Author: playlistTrack.Author, Title: playlistTrack.Title}
		resp.TrackList = append(resp.TrackList, track)
	}

	_ = json.NewEncoder(w).Encode(resp)
}

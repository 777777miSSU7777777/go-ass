package api

import (
	"github.com/gorilla/mux"

	"github.com/777777miSSU7777777/go-ass/middleware"
)

func NewAPIRouter(r *mux.Router, api API) {
	publicAPI := r.PathPrefix("/api").Subrouter()
	publicAPI.Use(middleware.JsonTypeMiddleware)
	publicAPI.Use(middleware.AllowCorsMiddleware)
	publicAPI.Methods("GET").Path("/audio/tracks").HandlerFunc(api.GetAllTracks)
	publicAPI.Methods("GET").Path("/audio/tracks/{id}").HandlerFunc(api.GetTrackByID)
	publicAPI.Methods("GET").Path("/audio/playlists").HandlerFunc(api.GetAllPlaylists)
	publicAPI.Methods("GET").Path("/audio/playlists/{id}").HandlerFunc(api.GetPlaylistByID)

	authorizedAPI := r.PathPrefix("/api").Subrouter()
	authorizedAPI.Use(JwtAuthMiddleware)
	authorizedAPI.Use(middleware.JsonTypeMiddleware)
	authorizedAPI.Use(middleware.AllowCorsMiddleware)
	authorizedAPI.Methods("POST", "OPTIONS").Path("/audio/tracks").HandlerFunc(api.AddTrack)
	authorizedAPI.Methods("PUT", "OPTIONS").Path("/audio/tracks/{id}").HandlerFunc(api.UpdateTrackByID)
	authorizedAPI.Methods("DELETE", "OPTIONS").Path("/audio/tracks/{id}").HandlerFunc(api.DeleteTrackByID)
	authorizedAPI.Methods("GET").Path("/audio/user-list").HandlerFunc(api.GetUserTrackList)
	authorizedAPI.Methods("PATCH", "OPTIONS").Path("/audio/user-list/add").HandlerFunc(api.AddTrackToUserTrackList)
	authorizedAPI.Methods("PATCH", "OPTIONS").Path("/audio/user-list/remove").HandlerFunc(api.RemoveTrackFromUserTrackList)
	authorizedAPI.Methods("POST", "OPTIONS").Path("/audio/playlists").HandlerFunc(api.CreateNewPlaylist)
	authorizedAPI.Methods("DELETE", "OPTIONS").Path("/audio/playlists/{id}").HandlerFunc(api.DeletePlaylistByID)
	authorizedAPI.Methods("GET").Path("/audio/user-playlists").HandlerFunc(api.GetUserPlaylists)
	authorizedAPI.Methods("PATCH", "OPTIONS").Path("/audio/playlists/{id}/add").HandlerFunc(api.AddTracksToPlaylist)
	authorizedAPI.Methods("PATCH", "OPTIONS").Path("/audio/playlists/{id}/remove").HandlerFunc(api.RemoveTracksFromPlaylist)
}

func NewAuthRouter(r *mux.Router, api API) {
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(middleware.AllowCorsMiddleware)
	s.Use(middleware.JsonTypeMiddleware)

	s.Methods("POST", "OPTIONS").Path("/signup").HandlerFunc(api.SignUp)
	s.Methods("POST", "OPTIONS").Path("/signin").HandlerFunc(api.SignIn)
	s.Methods("POST", "OPTIONS").Path("/refresh-token").HandlerFunc(api.RefreshToken)
	s.Methods("DELETE", "OPTIONS").Path("/signout").HandlerFunc(api.SignOut)
}

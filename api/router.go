package api

import (
	"github.com/gorilla/mux"

	"github.com/777777miSSU7777777/go-ass/middleware"
)

func NewAPIRouter(r *mux.Router, api API) {
	publicAPI := r.PathPrefix("/api").Subrouter()
	publicAPI.Use(middleware.JsonTypeMiddleware)
	publicAPI.Use(middleware.AllowCorsMiddleware)
	publicAPI.Methods("GET").Path("/audio").HandlerFunc(api.GetAudioList)
	publicAPI.Methods("GET").Path("/audio/{id}").HandlerFunc(api.GetAudioByID)
	publicAPI.Methods("GET").Path("/audio-playlists").HandlerFunc(api.GetAllAudioPlaylists)
	publicAPI.Methods("GET").Path("/audio-playlists/{id}").HandlerFunc(api.GetAudioPlaylistByID)

	authorizedAPI := r.PathPrefix("/api").Subrouter()
	authorizedAPI.Use(JwtAuthMiddleware)
	authorizedAPI.Use(middleware.JsonTypeMiddleware)
	authorizedAPI.Use(middleware.AllowCorsMiddleware)
	authorizedAPI.Methods("POST").Path("/audio").HandlerFunc(api.AddAudio)
	authorizedAPI.Methods("PUT").Path("/audio/{id}").HandlerFunc(api.UpdateAudioByID)
	authorizedAPI.Methods("DELETE", "OPTIONS").Path("/audio/{id}").HandlerFunc(api.DeleteAudioByID)
	authorizedAPI.Methods("GET").Path("/user-audio").HandlerFunc(api.GetUserAudioList)
	authorizedAPI.Methods("POST").Path("/user-audio").HandlerFunc(api.AddAudioToUserAudioList)
	authorizedAPI.Methods("DELETE").Path("/user-audio").HandlerFunc(api.DeleteAudioFromUserAudioList)
	authorizedAPI.Methods("POST").Path("/audio-playlists").HandlerFunc(api.CreateNewPlaylist)
	authorizedAPI.Methods("PATCH").Path("/audio-playlists/{id}/add").HandlerFunc(api.AddAudioListToPlaylist)
	authorizedAPI.Methods("PATCH").Path("/audio-playlists/{id}/delete").HandlerFunc(api.DeleteAudioListFromPlaylist)
}

func NewAuthRouter(r *mux.Router, api API) {
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(middleware.JsonTypeMiddleware)

	s.Methods("POST").Path("/signup").HandlerFunc(api.SignUp)
	s.Methods("POST").Path("/signin").HandlerFunc(api.SignIn)
	s.Methods("POST").Path("/refresh-token").HandlerFunc(api.RefreshToken)
	s.Methods("DELETE").Path("/signout").HandlerFunc(api.SignOut)
}

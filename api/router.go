package api

import (
	"github.com/gorilla/mux"

	"github.com/777777miSSU7777777/go-ass/middleware"
)

func NewAPIRouter(r *mux.Router, api API) {
	s := r.PathPrefix("/api").Subrouter()
	s.Use(JwtAuthMiddleware)
	s.Use(middleware.JsonTypeMiddleware)
	s.Use(middleware.AllowCorsMiddleware)

	s.Methods("POST").Path("/audio").HandlerFunc(api.AddAudio)
	s.Methods("GET").Path("/audio").HandlerFunc(api.GetAudioList)
	s.Methods("GET").Path("/audio/{id}").HandlerFunc(api.GetAudioByID)
	s.Methods("PUT").Path("/audio/{id}").HandlerFunc(api.UpdateAudioByID)
	s.Methods("DELETE", "OPTIONS").Path("/audio/{id}").HandlerFunc(api.DeleteAudioByID)
}

func NewAuthRouter(r *mux.Router, api API) {
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(middleware.JsonTypeMiddleware)

	s.Methods("POST").Path("/signup").HandlerFunc(api.SignUp)
	s.Methods("POST").Path("/signin").HandlerFunc(api.SignIn)
	s.Methods("POST").Path("/refresh-token").HandlerFunc(api.RefreshToken)
	s.Methods("DELETE").Path("/signout").HandlerFunc(api.SignOut)
}

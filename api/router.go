package api

import (
	"github.com/gorilla/mux"

	"github.com/777777miSSU7777777/go-ass/middleware"
)

func NewAPIRouter(r *mux.Router, api API) {
	s := r.PathPrefix("/api").Subrouter()
	s.Use(middleware.JsonTypeMiddleware)
	s.Use(middleware.NoCorsMiddleware)

	s.Methods("POST").Path("/audio").HandlerFunc(api.AddAudio)
	s.Methods("GET").Path("/audio").HandlerFunc(api.GetAudioList)
	s.Methods("GET").Path("/audio/{id:[0-9]+}").HandlerFunc(api.GetAudioByID)
	s.Methods("PUT").Path("/audio/{id:[0-9]+}").HandlerFunc(api.UpdateAudioByID)
	s.Methods("DELETE").Path("/audio/{id:[0-9]+}").HandlerFunc(api.DeleteAudioByID)
}

package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPIRouter(r *mux.Router, api API) {
	s := r.PathPrefix("/api").Subrouter()
	s.Use(jsonTypeMiddleware)

	s.Methods("POST").Path("/audio").HandlerFunc(api.AddAudio)
	s.Methods("GET").Path("/audio").HandlerFunc(api.GetAllAudio)
	s.Methods("GET").Path("/audio/{id:[0-9]+}").HandlerFunc(api.GetAudioByID)
	s.Methods("GET").Path("/audio/{key}").HandlerFunc(api.GetAudioByKey)
	s.Methods("PUT").Path("/audio/{id:[0-9]+}").HandlerFunc(api.UpdateAudioByID)
	s.Methods("DELETE").Path("/audio/{id:[0-9]+}").HandlerFunc(api.DeleteAudioByID)
}

func jsonTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

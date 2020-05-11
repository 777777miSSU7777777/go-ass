package stream

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/777777miSSU7777777/go-ass/middleware"
)

func NewStreamRouter(r *mux.Router, api StreamAPI) http.Handler {
	s := r.PathPrefix("/media").Subrouter()
	s.Use(middleware.AllowCorsMiddleware)

	s.Methods("GET").Path("/{id}/stream/").HandlerFunc(api.Stream)
	s.Methods("GET").Path("/{id}/stream/{quality_manifest}").HandlerFunc(api.Stream)
	s.Methods("GET").Path("/{id}/stream/{seg:seg[0-9]+_(?:64k|?:96k|?:128k|?:192k).ts}").HandlerFunc(api.Stream)
	s.Methods("GET").Path("/{id}/download").HandlerFunc(api.Download)

	return r
}

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
	s.Methods("GET").Path("/{id}/stream/{seg:seg[0-9]+.ts}").HandlerFunc(api.Stream)
	s.Methods("GET").Path("/{id}/download").HandlerFunc(api.Download)

	return r
}

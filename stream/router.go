package stream

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/777777miSSU7777777/go-ass/middleware"
)

func NewStreamRouter(r *mux.Router, api StreamAPI) http.Handler {
	s := r.PathPrefix("/media").Subrouter()
	s.Use(middleware.NoCorsMiddleware)

	s.Methods("GET").Path("/{id:[0-9]+}/stream/").HandlerFunc(api.Stream)
	s.Methods("GET").Path("/{id:[0-9]+}/stream/{seg:seg[0-9]+.ts}").HandlerFunc(api.Stream)
	s.Methods("GET").Path("/{id:[0-9]+}/download").HandlerFunc(api.Download)

	return r
}

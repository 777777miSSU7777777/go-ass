package stream

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewStreamRouter(r *mux.Router, api StreamAPI) http.Handler {
	s := r.PathPrefix("/media").Subrouter()

	s.Methods("GET").Path("/{id:[0-9]+}/stream/").HandlerFunc(api.Stream)
	s.Methods("GET").Path("/{id:[0-9]+}/stream/{seg:seg[0-9]+.ts}").HandlerFunc(api.Stream)

	return r
}

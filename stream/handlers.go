package stream

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type StreamAPI struct {
	m MediaManager
}

func NewStreamAPI(m MediaManager) StreamAPI {
	return StreamAPI{m}
}

func (a StreamAPI) Stream(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	seg, ok := vars["seg"]
	if !ok {
		a.m.ServeM3u8(w, r, id)
	} else {
		a.m.ServeTs(w, r, id, seg)
	}
}

func (a StreamAPI) Download(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	a.m.ServerMp3(w, r, id)
}

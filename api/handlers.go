package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/777777miSSU7777777/go-ass/model"
	"github.com/777777miSSU7777777/go-ass/repository"
	"github.com/777777miSSU7777777/go-ass/service"
)

type ErrorResponse struct {
	Type  string `json:"type"`
	Error string `json:"error"`
}

var BodyParseError = "BODY PARSE ERROR"
var IDParseError = "ID PARSE ERROR"
var ValidationError = "VALIDATION ERROR"
var NotFoundError = "NOT FOUND ERROR"
var ServiceError = "SERVICE ERROR"
var QueryStringError = "QUERY STRING ERROR"
var InternalServerError = "INTERNAL SERVER ERROR"

func writeError(w http.ResponseWriter, statusCode int, errType string, err error) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Type: errType, Error: err.Error()})
}

type API struct {
	svc service.Service
	m   UploadManager
}

func NewApi(svc service.Service, m UploadManager) API {
	return API{svc, m}
}

func (a API) AddAudio(w http.ResponseWriter, r *http.Request) {
	var req AddAudioRequest

	req.Author, req.Title = r.FormValue("author"), r.FormValue("title")

	lastID, err := a.svc.GetLastAudioID()
	if err != nil {
		writeError(w, 500, InternalServerError, fmt.Errorf("internal server error: %v", err))
		return
	}

	err = a.m.Upload(w, r, lastID+1)
	if err != nil {
		return
	}

	resp, err := a.svc.AddAudio(req.Author, req.Title)
	if err != nil {
		if err.Error() == model.AudioAuthorEmpty.Error() || err.Error() == model.AudioTitleEmpty.Error() {
			writeError(w, 400, ValidationError, err)
			fmt.Println(err)
		} else {
			writeError(w, 400, ServiceError, err)
			fmt.Println(err)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(AddAudioResponse{resp.ID, resp.Author, resp.Title})
}

func (a API) GetAllAudio(w http.ResponseWriter, r *http.Request) {
	resp, err := a.svc.GetAllAudio()
	if err != nil {
		if err.Error() == repository.AudioNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(GetAllAudioResponse{resp})
}

func (a API) GetAudioByID(w http.ResponseWriter, r *http.Request) {
	var req GetAudioByIDRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, 400, IDParseError, fmt.Errorf("error while parsing id: %v", err))
		return
	}
	req.ID = id

	resp, err := a.svc.GetAudioByID(req.ID)
	if err != nil {
		if err.Error() == repository.AudioNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(GetAudioByIDResponse{resp.ID, resp.Author, resp.Title})
}

func (a API) GetAudioByKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		writeError(w, 400, QueryStringError, fmt.Errorf("key for search not found error"))
		return
	}
	audio := make([]model.Audio, 0)

	authorResp, err := a.svc.GetAudioByAuthor(key)
	if err != nil {
		if err.Error() == repository.AudioNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}
	audio = append(audio, authorResp...)

	titleResp, err := a.svc.GetAudioByTitle(key)
	if err != nil {
		if err.Error() == repository.AudioNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	audio = append(audio, titleResp...)

	if len(audio) == 0 {
		writeError(w, 404, NotFoundError, err)
		return
	}

	_ = json.NewEncoder(w).Encode(GetAudioByKeyResponse{audio})
}

func (a API) UpdateAudioByID(w http.ResponseWriter, r *http.Request) {
	var req UpdateAudioByIDRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, 400, IDParseError, fmt.Errorf("error while parsing id: %v", err))
		return
	}
	req.ID = id

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	resp, err := a.svc.UpdateAudioByID(req.ID, req.Author, req.Title)
	if err != nil {
		if err.Error() == repository.AudioNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(UpdateAudioByIDResponse{resp.ID, resp.Author, resp.Title})
}

func (a API) DeleteAudioByID(w http.ResponseWriter, r *http.Request) {
	var req DeleteAudioByIDRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, 400, IDParseError, fmt.Errorf("error while parsing id: %v", err))
		return
	}
	req.ID = id

	err = a.svc.DeleteAudioByID(req.ID)
	if err != nil {
		if err.Error() == repository.AudioNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(DeleteAudioByIDResponse{})
}

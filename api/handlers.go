package api

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	m   FileManager
}

func NewApi(svc service.Service, m FileManager) API {
	return API{svc, m}
}

func (a API) AddAudio(w http.ResponseWriter, r *http.Request) {
	var req AddAudioRequest

	req.Author, req.Title = r.FormValue("author"), r.FormValue("title")

	resp, err := a.svc.AddAudio(req.Author, req.Title, r.Context().Value("userID").(string))
	if err != nil {
		if err.Error() == model.AudioAuthorEmpty.Error() || err.Error() == model.AudioTitleEmpty.Error() {
			writeError(w, 400, ValidationError, err)
			fmt.Println(err)
		} else {
			writeError(w, 400, ServiceError, err)
			fmt.Println(err)
		}
		_ = a.m.Delete(w, resp.ID)
		return
	} else {
		err = a.m.Upload(w, r, resp.ID)
		if err != nil {
			_ = a.svc.DeleteAudioByID(resp.ID)
			_ = a.m.Delete(w, resp.ID)
			return
		}
	}

	_ = json.NewEncoder(w).Encode(AddAudioResponse{resp.ID, resp.Author, resp.Title})
}

func (a API) GetAudioList(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	var resp []model.Audio
	var err error
	if key != "" {
		resp, err = a.svc.GetAudioByKey(key)
		if err != nil {
			if err.Error() == repository.AudioNotFoundError.Error() {
				writeError(w, 400, NotFoundError, err)
			} else {
				writeError(w, 400, ServiceError, err)
			}
			return
		}
	} else {
		resp, err = a.svc.GetAllAudio()
		if err != nil {
			if err.Error() == repository.AudioNotFoundError.Error() {
				writeError(w, 404, NotFoundError, err)
			} else {
				writeError(w, 400, ServiceError, err)
			}
			return
		}
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (a API) GetAudioByID(w http.ResponseWriter, r *http.Request) {
	var req GetAudioByIDRequest
	vars := mux.Vars(r)
	req.ID = vars["id"]

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

func (a API) UpdateAudioByID(w http.ResponseWriter, r *http.Request) {
	var req UpdateAudioByIDRequest
	vars := mux.Vars(r)
	req.ID = vars["id"]

	err := json.NewDecoder(r.Body).Decode(&req)
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
	req.ID = vars["id"]

	err := a.m.Delete(w, req.ID)
	if err != nil {
		return
	}

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

func (a API) SignUp(w http.ResponseWriter, r *http.Request) {
	var req SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	err = a.svc.SignUp(req.Email, req.Name, req.Password)
	if err != nil {
		if err.Error() == repository.UserNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(SignUpResponse{})
}

func (a API) SignIn(w http.ResponseWriter, r *http.Request) {
	var req SignInRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	accessToken, refreshToken, err := a.svc.SignIn(req.Email, req.Password)
	if err != nil {
		if err.Error() == repository.UserNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
	}

	_ = json.NewEncoder(w).Encode(SignInResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

func (a API) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshTokenRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	accessToken, refreshToken, err := a.svc.RefreshToken(req.RefreshToken)
	if err != nil {
		if err.Error() == repository.RefreshTokenNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
	}

	_ = json.NewEncoder(w).Encode(SignInResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

func (a API) SignOut(w http.ResponseWriter, r *http.Request) {
	var req SignOutRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	err = a.svc.SignOut(req.RefreshToken)
	if err != nil {
		if err.Error() == repository.RefreshTokenNotFoundError.Error() {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, err)
		}
	}

	_ = json.NewEncoder(w).Encode(SignOutResponse{})
}

package api

type AddAudioRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

type GetAudioByIDRequest struct {
	ID int64
}

type UpdateAudioByIDRequest struct {
	ID     int64
	Author string `json:"author"`
	Title  string `json:"title"`
}

type DeleteAudioByIDRequest struct {
	ID int64
}

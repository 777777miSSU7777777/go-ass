package api

type AddAudioRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

type GetAudioByIDRequest struct {
	ID string
}

type UpdateAudioByIDRequest struct {
	ID     string
	Author string `json:"author"`
	Title  string `json:"title"`
}

type DeleteAudioByIDRequest struct {
	ID string
}

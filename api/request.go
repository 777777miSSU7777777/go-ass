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

type SignUpRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct{
	RefreshToken string `json:"refresh_token"`
}

type SignOutRequest RefreshTokenRequest
package dto

type AuthResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

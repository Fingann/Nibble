package user

type UserRegistrationResponse struct {
	ID uint `json:"id"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

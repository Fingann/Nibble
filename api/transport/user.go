package user

type RegistrationRequest struct {
	Username string `json:"username" binding:"required" `
	Email    string `json:"email" binding:"required,email" `
	Password string `json:"password" binding:"required" `
}

type RegistrationResponse struct {
	ID    uint   `json:"id" `
	Error string `json:"error,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" `
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Error string `json:"error,omitempty"`
}

type UserDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}

type UserDeleteResponse struct {
	ID    uint   `json:"id"`
	Error string `json:"error,omitempty"`
}

type UserRetrieveRequest struct {
	Id uint `json:"id" form:"Id" binding:"required"`
}

type UserRetrieveResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Error    string `json:"error,omitempty"`
}

type UserUpdateRequest struct {
	Id       uint   `json:"id" form:"Id" binding:"required"`
	Username string `json:"username"`
	Email    string `json:"email" binding:"email"`
}

type UserUpdateResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Error    string `json:"error,omitempty"`
}

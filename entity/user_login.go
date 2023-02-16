package entity

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	// success：0，other：not 0
	StatusCode int8 `json:"status_code"`
	// login message
	StatusMsg string `json:"status_msg,omitempty"`
	// ID of user
	UserId uint `json:"user_id,omitempty"`
	// token according username and password
	Token string `json:"token,omitempty"`
}

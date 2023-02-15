package entity

type user_login_request struct {
	// 用户名
	Username string `json:"username"`
	// 密码
	Password string `json:"password"`
}

type user_login_response struct {
	// 状态码，0表示成功，其它值表示失败
	StatusCode string `json:"status_code"`
	// 登陆成功状态信息
	StatusMsg string `json:"status_msg omitempty"`
	// 用户ID
	UserId string `json:"user_id omitempty"`
	// 鉴权token
	Token string `json:"token omitempty"`
}

package dto

// 用户登录表单
type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// 登录响应
type LoginResponse struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}

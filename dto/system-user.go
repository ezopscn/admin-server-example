package dto

// 基础用户信息
type BaseUserInfo struct {
	Id       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	CNName   string `json:"cn_name" form:"cn_name"`
	ENName   string `json:"en_name" form:"en_name"`
}

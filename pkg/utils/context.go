package utils

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 从 Context 中获取 String 属性
func ExtractStringFromContext(ctx *gin.Context, keyword string) (value string, err error) {
	claims := jwt.ExtractClaims(ctx)
	value, ok := claims[keyword].(string)
	if !ok || value == "" {
		err = fmt.Errorf("从 Context 中获取字符串失败")
	}
	return
}

// 从 Context 中获取 Username
func GetUsernameFromContext(ctx *gin.Context) (username string, err error) {
	username, err = ExtractStringFromContext(ctx, "Username")
	return
}

// 从 Context 中获取 Phone
func GetPhoneFromContext(ctx *gin.Context) (phone string, err error) {
	phone, err = ExtractStringFromContext(ctx, "Phone")
	return
}

// 从 Context 中获取 Email
func GetEmailFromContext(ctx *gin.Context) (email string, err error) {
	email, err = ExtractStringFromContext(ctx, "Email")
	return
}

// 从 Context 中获取 Uint 属性
func ExtractUintFromContext(ctx *gin.Context, keyword string) (value uint, err error) {
	claims := jwt.ExtractClaims(ctx)
	v, _ := claims[keyword].(float64) // 注意客户端请求过来的 JSON 会变成 float64 类型
	if v == 0 {
		return value, fmt.Errorf("获取请求用户的%s失败", keyword)
	}
	value = uint(v)
	return
}

// 从 Context 中获取 RoleId
func GetRoleIdFromContext(ctx *gin.Context) (roleId uint, err error) {
	roleId, err = ExtractUintFromContext(ctx, "RoleId")
	return
}

// 从 Context 中获取 Username 和 RoleId
func GetUsernameAndRoleIdFromContext(ctx *gin.Context) (username string, roleId uint, err error) {
	username, err = ExtractStringFromContext(ctx, "Username")
	if err != nil {
		return
	}
	roleId, err = ExtractUintFromContext(ctx, "RoleId")
	return
}

package utils

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"server/common"
)

// 密码加密
func CryptoPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

// 密码验证
func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 生成 TOTP 双因子认证信息
func GenerateTOTPKey(account string) (key *otp.Key, err error) {
	key, err = totp.Generate(totp.GenerateOpts{
		Issuer:      common.APPName,    // 应用名称
		AccountName: account,           // 用户名称
		Period:      30,                // 密码的有效时间，一般 30 秒
		Digits:      otp.DigitsSix,     // 生成的密码长度，一般为 6 位
		Algorithm:   otp.AlgorithmSHA1, // 用于 HMAC 签名的算法，默认是 SHA1
	})
	return
}

// 验证 TOTP 双因子认证验证码是否正确
func ValidateTOTPCode(code string, secret string) bool {
	if valid := totp.Validate(code, secret); valid {
		return true
	}
	return false
}

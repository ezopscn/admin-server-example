package service

import (
	"github.com/bytedance/sonic"
	"server/common"
	"server/model"
	"server/pkg/gedis"
)

// 根据用户名查询用户信息
func GetBaseUserInfoByUsername(username string) (user model.User, err error) {
	// 查询 Redis 中是否存在该数据
	var cache = gedis.NewOperation()
	userInfoKey := common.GenerateRedisKey(common.RKP.UserInfo, username)
	dataStr := cache.GetString(userInfoKey).Unwrap()
	if dataStr != "" {
		err = sonic.Unmarshal([]byte(dataStr), &user)
		if err == nil {
			return
		}
	}

	// 基础关联信息, 使用 clause.Associations 可以预加载所有字段
	err = common.DB.
		Preload("ManageDepartments").
		Preload("Departments").
		Preload("Job").
		Preload("Role").
		Preload("OfficeProvince").
		Preload("OfficeCity").
		Preload("NativeProvince").
		Preload("NativeCity").
		Where("username = ?", username).First(&user).Error
	if err != nil {
		return
	}
	jsonData, _ := sonic.Marshal(user)
	cache.Set(userInfoKey, jsonData, gedis.WithExpire(common.UserInfoExpireTime))
	return
}

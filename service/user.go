package service

import (
	"mallapi/global"
	"mallapi/models"
)

type WebUserService struct {
}

// 用户登录信息验证
func (u *WebUserService) Login(param models.WebLoginParam) uint64 {
	var user models.User
	global.Db.Where("username = ? and password = ?", param.Username, param.Password).First(&user)
	return user.Id
}

// 用户注册
func (u *WebUserService) SignUp(param models.User) bool {
	global.Db.Create(&param)
	return true
}
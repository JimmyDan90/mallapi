package models

// User 数据库，用户数据映射模型
type User struct {
	Id uint64 `gorm:"primaryKey"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Phone string `gorm:"phone"`
	Status uint `gorm:"status"`
}

// WebLoginParam 后台管理前端，用户登录参数模型
type WebLoginParam struct {
	Username string `json:"username"` // 用户名称
	Password string `json:"password"` // 登录密码
	Phone string `json:"phone"` // 手机号
	CaptchaId string `gorm:"captchaId"` // 验证码接口返回的字符串
	CaptchaValue string `gorm:"captchaValue"` // 验证码
}


// 后台管理前端，用户信息传输模型
type WebUserInfo struct {
	Uid uint64 `json:"uid"`
	Token string `json:"token"`
}

// 微信小程序，用户登录凭证校验模型
type AppCode2Session struct {
	Code string
	AppId string
	AppSecret string
}

// 微信小程序，凭证校验后返回的JSON数据包模型
type AppCode2SessionJson struct {
	OpenId string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId string `json:"unionid"`
	ErrCode uint `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

package api

import (
	"github.com/gin-gonic/gin"
	"mallapi/common"
	"mallapi/models"
	"mallapi/response"
	"mallapi/service"
)

var user service.WebUserService

// WebGetCaptcha 后台管理系统，获取验证码
func WebGetCaptcha(c *gin.Context) {
	id, b64s, _ := common.GenerateCaptcha()
	data := map[string]interface{}{"captchaId": id, "captchaImg": b64s}
	response.Success("操作成功", data, c)
}

// WebUserLogin web后台系统登录接口
// @Summary web后台系统登录接口
// @Description 通过token校验通过后才可以登录
// @Accept application/json
// @Produce application/json
// @Param object body models.WebLoginParam true "用户登录参数"
// @Success 200 {object} models.WebUserInfo
// @Router /login [post]
func WebUserLogin(c *gin.Context) {
	var param models.WebLoginParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	// 检查验证码
	if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue) {
		response.Failed("验证码错误", c)
		return
	}
	// 生成token
	uid := user.Login(param)
	if uid > 0 {
		token, _ := common.GenerateToken(param.Username)
		userInfo := models.WebUserInfo{
			Uid: uid,
			Token: token,
		}
		response.Success("登录成功", userInfo, c)
		return
	}
	response.Failed("用户名或密码错误", c)
}

// WebUserSignUp 后台管理系统，用户注册
func WebUserSignUp(c *gin.Context) {
	var param models.WebLoginParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	// 检查验证码
	if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue) {
		response.Failed("验证码错误", c)
		return
	}
	// 用户注册
	var signInfo = models.User{
		Username: param.Username,
		Password: param.Password,
		Phone: param.Phone,
	}
	success := user.SignUp(signInfo)
	if success {
		response.Success("注册成功", nil, c)
		return
	}
}
package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mallapi/api"
	"mallapi/global"
)

func Router() {
	engine := gin.Default()

	// 静态资源请求映射
	engine.Static("/image", global.Config.Upload.SavePath)
	// 后台管理员前端接口
	web := engine.Group("/web")
	{
		// 用户登录API
		web.GET("/captcha", api.WebGetCaptcha)
		web.POST("/sign", api.WebUserSignUp)
		web.POST("/login", api.WebUserLogin)
		// 开启JWT认证，以下接口需要认证成功才能访问
	}
	// 启动、监听接口
	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("server start error: %s", err)
	}
}

package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mallapi/api"
	"mallapi/global"
	"mallapi/middleware"
)

func Router() {
	engine := gin.Default()

	// 静态资源请求映射
	engine.Static("/image", global.Config.Upload.SavePath)
	// 后台管理员前端接口
	web := engine.Group("/web/api/v1")
	{
		// 用户登录API
		web.GET("/captcha", api.WebGetCaptcha)
		web.POST("/sign", api.WebUserSignUp)
		web.POST("/login", api.WebUserLogin)
		// 开启JWT认证，以下接口需要认证成功才能访问
		web.Use(middleware.JwtAuth())

		// 类目管理API
		web.POST("/category/create", api.WebCreateCategory)
		web.DELETE("/category/delete", api.WebDeleteCategory)
		web.PUT("/category/update", api.WebUpdateCategory)
		web.GET("/category/list", api.WebGetCategoryList)
		web.GET("/category/option", api.WebGetCategoryOption)
	}
	// 启动、监听接口
	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("server start error: %s", err)
	}
}

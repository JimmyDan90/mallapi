package middleware

import (
	"github.com/gin-gonic/gin"
	"mallapi/common"
	"mallapi/response"
)

// JwtAuth JWT认证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Failed("未登录或非法访问", c)
			c.Abort()
			return
		}
		_, err := common.VerifyToken(token)
		if err != nil {
			response.Failed("token校验失败，请重新登录", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
package middlewares

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/3/2 20:41
 * @Desc:
 */

// GuestJWT 强制使用游客身份访问
func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.GetHeader("Authorization")) > 0 {
			// 解析 token 成功，说明登录成功了
			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				response.Unauthorized(c, "请使用游客身份访问")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

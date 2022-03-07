package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gohub/pkg/response"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/3/7 13:29
 * @Desc:
 */

// ForceUA 中间件，强制请求必须附带 User-Agent 标头
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 User-Agent 标头信息
		if len(c.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(c, errors.New("User-Agent 标头未找到"), "请求必须附带 User-Agent 标头")
			return
		}

		c.Next()
	}
}

package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/routes"
	"net/http"
	"strings"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/2/28 18:18
 * @Desc:
 */

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(router)

	//  注册 API 路由
	routes.RegisterAPIRoutes(router)

	//  配置 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(r *gin.Engine) {
	// 处理 404 请求
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}

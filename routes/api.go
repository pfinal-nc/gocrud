package routes

import (
	"github.com/gin-gonic/gin"
	controllers "gohub/app/http/controllers/api/v1"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"
	"gohub/pkg/config"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/2/28 19:27
 * @Desc:
 */

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断邮箱是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 通过手机号码注册
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			// 通过邮箱注册
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)

			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password", lgc.LoginByPassword)

			// 刷新 Token
			authGroup.POST("/login/refresh-token", lgc.RefreshToken)

			uc := new(controllers.UsersController)
			// 获取当前用户
			v1.GET("/user", middlewares.AuthJwt(), uc.CurrentUser)
			usersGroup := v1.Group("/users")
			{
				usersGroup.GET("", uc.Index)
				usersGroup.PUT("", middlewares.AuthJwt(), uc.UpdateProfile)
				usersGroup.PUT("/avatar", middlewares.AuthJwt(), uc.UpdateAvatar)
			}

			cgc := new(controllers.CategoriesController)
			cgcGroup := v1.Group("/categories")
			{
				cgcGroup.GET("", cgc.Index)
				cgcGroup.POST("", middlewares.AuthJwt(), cgc.Store)
				cgcGroup.PUT("/:id", middlewares.AuthJwt(), cgc.Update)
				cgcGroup.DELETE("/:id", middlewares.AuthJwt(), cgc.Delete)
			}
			tpc := new(controllers.TopicsController)
			tpcGroup := v1.Group("/topics")
			{
				tpcGroup.GET("", tpc.Index)
				tpcGroup.POST("", middlewares.AuthJwt(), tpc.Store)
				tpcGroup.PUT("/:id", middlewares.AuthJwt(), tpc.Update)
				tpcGroup.DELETE("/:id", middlewares.AuthJwt(), tpc.Delete)
			}
		}
	}
}

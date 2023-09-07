package auth

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/3/1 17:04
 * @Desc:
 */

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/app/models/user"
	"server/pkg/config"
	"server/pkg/logger"
)

// Attempt 尝试登录
func Attempt(name string, password string) (user.User, error) {
	userModel := user.GetByMulti(name)
	if userModel.UserId == 0 {
		return user.User{}, errors.New("账号不存在")
	}
	// TODO 处理超级用户
	fmt.Println(config.Get("supper"))

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.UserId == 0 {
		return user.User{}, errors.New("手机号未注册")
	}

	return userModel, nil
}

// CurrentUser 从 gin.context 中获取当前登录的用户

func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}

	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}

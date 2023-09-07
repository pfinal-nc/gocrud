package factories

import (
	"github.com/bxcodec/faker/v3"
	"server/app/models/user"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/3/4 11:30
 * @Desc:
 */

func MakeUsers(times int) []user.User {
	var objs []user.User
	// 设置唯一值
	faker.SetGenerateUniqueValues(true)
	for i := 0; i < times; i++ {
		model := user.User{
			Name: faker.Username(),
			//Email:    faker.Email(),
			//Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}
		objs = append(objs, model)
	}
	return objs
}

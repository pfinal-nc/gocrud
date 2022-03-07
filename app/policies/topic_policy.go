package policies

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/topic"
	"gohub/pkg/auth"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/3/7 10:55
 * @Desc:
 */

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}

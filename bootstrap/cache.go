package bootstrap

import (
	"fmt"
	"gohub/pkg/cache"
	"gohub/pkg/config"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/3/7 13:11
 * @Desc:
 */

// SetupCache 缓存
func SetupCache() {

	// 初始化缓存专用的 redis client, 使用专属缓存 DB
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)

	cache.InitWithCacheStore(rds)
}

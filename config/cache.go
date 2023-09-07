package config

import "server/pkg/config"

func init() {

	config.Add("cache", func() map[string]interface{} {
		return map[string]interface{}{
			// 过期时间，单位是分钟，一般不超过两个小时
			"expire_time": config.Env("CACHE_EXPIRE_TIME", 3600),
		}
	})
}

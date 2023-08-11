package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"server/pkg/cache"
	"server/pkg/console"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

// forget 命令的选项
var cacheKey string

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "Clear cache",
	Run:   runCacheClear,
}

func init() {
	// 注册 cache 命令的子命令
	CmdCache.AddCommand(CmdCacheClear, CmdCacheForget)
	// 设置 cache forget 命令的选项
	CmdCacheForget.Flags().StringVarP(&cacheKey, "key", "k", "", "KEY of the cache")
	CmdCacheForget.MarkFlagRequired("key")
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared.")
}

var CmdCacheForget = &cobra.Command{
	Use:   "forget",
	Short: "Delete redis key, example: cache forget cache-key",
	Run:   runCacheForget,
}

func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("Cache key [%s] deleted.", cacheKey))
}

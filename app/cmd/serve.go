package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"server/bootstrap"
	"server/pkg/config"
	"server/pkg/console"
	"server/pkg/logger"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/3/3 16:42
 * @Desc:
 */

// CmdServe represents the available web sub-command.
var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode("release")

	// gin 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务器
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}

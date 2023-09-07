package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"server/pkg/console"
)

var CmdMakeMiddleware = &cobra.Command{
	Use:   "middleware",
	Short: "Create middleware file, exmaple: make middleware ",
	Run:   runMakeMiddleware,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeMiddleware(cmd *cobra.Command, args []string) {

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/http/middlewares/%s.go", args[0])
	model := makeModelFromBaseString(args[0])
	// 基于模板创建文件（做好变量替换）
	createFileFromStubBase(filePath, "middleware", model)

	console.Success("the middleware file created")
}

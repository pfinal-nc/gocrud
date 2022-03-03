package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/3/3 19:39
 * @Desc:
 */

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file, example make request user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeRequest(cmd *cobra.Command, args []string) {

	// 格式化模型名称，返回一个 Model 对象
	model := makeModelFromString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/requests/%s_request.go", model.PackageName)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "request", model)
}

package tools

import "github.com/spf13/cobra"

var CmdTools = &cobra.Command{
	Use:   "tools",
	Short: "Generate file and code",
}

func init() {
	// 注册 make 的子命令
	CmdTools.AddCommand(
		CmdGeneratePass,
	)
}

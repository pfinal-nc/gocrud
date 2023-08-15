package tools

import (
	"github.com/spf13/cobra"
	"server/pkg/console"
	"server/pkg/hash"
)

var CmdGeneratePass = &cobra.Command{
	Use:   "pwd",
	Short: "Generates a password set by the system",
	Run:   runGeneratePass,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runGeneratePass(cmd *cobra.Command, args []string) {
	console.Success("command name:" + hash.BcryptHash(args[0]))
}

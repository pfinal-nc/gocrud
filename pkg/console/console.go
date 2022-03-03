package console

import (
	"fmt"
	"github.com/mgutz/ansi"
	"os"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/3/2 20:59
 * @Desc:
 */

// Success 打印一条成功消息，绿色输出
func Success(msg string) {
	colorOut(msg, "green")
}

// Error 打印一条报错信息 红色输出

// colorOut 内部使用，设置高亮颜色
func colorOut(message, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(message, color))
}

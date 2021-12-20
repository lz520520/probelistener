package log

import (
	"fmt"
	"runtime"
)

func isLin() bool {
	if runtime.GOOS == "linux" {
		return true
	} else {
		return false
	}
}
func LogInfo(msg string) {
	// 绿色
	fmt.Println("\033[32m [+] " + msg + " \033[0m")
}
func LogDebug(msg string) {
	// 蓝色
	fmt.Println("\033[34m [+] " + msg + " \033[0m")
}

func LogError(msg string) {
	// 红色
	fmt.Println("\033[31m [-] " + msg + " \033[0m")
}

func LogSuccess(msg string) {
	// 黄底红字
	fmt.Println("\033[43:32m [+] " + msg + " \033[0m")

}

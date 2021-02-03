package Os

import (
	"runtime"
)

// 判断是否为Linux系统
func (*Os)IsLinux() bool {
	return "linux" == runtime.GOOS
}

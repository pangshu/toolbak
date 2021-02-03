package Os

import (
	"runtime"
)

// 判断是否为Windows系统
func (*Os)IsWindows() bool {
	return "windows" == runtime.GOOS
}

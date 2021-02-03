package Os

import (
	"runtime"
)

// 判断是否为Mac系统
func (*Os)IsMac() bool {
	return "darwin" == runtime.GOOS
}

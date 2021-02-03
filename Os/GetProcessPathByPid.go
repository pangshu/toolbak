package Os

import (
	"fmt"
	"os"
)

func (*Os)GetProcessPathByPid(pid int) string {
	exe := fmt.Sprintf("/proc/%d/exe", pid)
	path, _ := os.Readlink(exe)
	return path
}

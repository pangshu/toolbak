package Os

import (
	"os"
	"syscall"
)

// IsProcessExists 进程是否存在.
func (*Os)IsProcessExists(pid int) (res bool) {
	process, err := os.FindProcess(pid)
	if err == nil {
		if err = process.Signal(os.Signal(syscall.Signal(0))); err == nil {
			res = true
		}
	}

	return
}
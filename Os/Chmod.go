package Os

import "os"

// Chmod 改变文件模式.
func (*Os)Chmod(filename string, mode os.FileMode) bool {
	return os.Chmod(filename, mode) == nil
}
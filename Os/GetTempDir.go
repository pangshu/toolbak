package Os

import "os"

// GetTempDir 返回用于临时文件的目录.
func (*Os)GetTempDir() string {
	return os.TempDir()
}

package File

import "os"

// GetFileMode 获取路径的权限模式.
func (*File)GetFileMode(filePath string) (os.FileMode, error) {
	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Mode(), nil
}

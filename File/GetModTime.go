package File

import "os"

// GetModTime 获取文件的修改时间戳,秒.
func (*File)GetModTime(filePath string) (res int64) {
	fileInfo, err := os.Stat(filePath)
	if err == nil {
		res = fileInfo.ModTime().Unix()
	}
	return
}

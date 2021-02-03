package File

import "os"

// FileSize 获取文件大小(bytes字节),注意:文件不存在或无法访问返回-1 .
func (*File)FileSize(filePath string) int64 {
	f, err := os.Stat(filePath)
	if nil != err {
		return -1
	}
	return f.Size()
}

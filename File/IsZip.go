package File

import (
	"bytes"
	"os"
)

// IsZip 是否zip文件.
func (toolFile *File)IsZip(fpath string) bool {
	ext := toolFile.Ext(fpath)
	if ext != "zip" {
		return false
	}

	f, err := os.Open(fpath)
	if err != nil {
		return false
	}
	defer func() {
		_ = f.Close()
	}()

	buf := make([]byte, 4)
	n, err := f.Read(buf)

	return err == nil && n == 4 && bytes.Equal(buf, []byte("PK\x03\x04"))
}

package File

import (
	"os"
	"path/filepath"
)

// ChmodBatch 批量改变路径权限模式(包括子目录和所属文件).filemode为文件权限模式,dirmode为目录权限模式.
func ChmodBatch(fpath string, filemode, dirmode os.FileMode) (res bool) {
	var err error
	err = filepath.Walk(fpath, func(fpath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			err = os.Chmod(fpath, dirmode)
		} else {
			err = os.Chmod(fpath, filemode)
		}

		return err
	})

	if err == nil {
		res = true
	}

	return
}

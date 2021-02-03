package Os

import "os"

// Chown 改变文件的所有者.
func (*Os)Chown(filename string, uid, gid int) bool {
	return os.Chown(filename, uid, gid) == nil
}

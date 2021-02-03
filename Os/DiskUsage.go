package Os

import "syscall"

// DiskUsage 获取磁盘/目录使用情况,单位字节.参数path为目录.
// used为已用,
// free为空闲,
// total为总数.
func (*Os)DiskUsage(path string) (used, free, total uint64) {
	fs := &syscall.Statfs_t{}
	err := syscall.Statfs(path, fs)
	if err == nil {
		total = fs.Blocks * uint64(fs.Bsize)
		free = fs.Bfree * uint64(fs.Bsize)
		used = total - free
	}

	return
}

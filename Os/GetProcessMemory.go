package Os

import "runtime"

// 获取当前go程序的内存使用,返回字节数.
func (*Os)GetProcessMemory() uint64 {
	stat := new(runtime.MemStats)
	runtime.ReadMemStats(stat)
	return stat.Alloc
}

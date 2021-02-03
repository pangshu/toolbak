package Os

import (
	"os"
	"runtime"
	"time"
)

// SystemInfo 系统信息
type RspSystem struct {
	ServerName   string  `json:"server_name"`    //服务器名称
	SystemOs     string  `json:"system_os"`      //操作系统名称
	Runtime      int64   `json:"run_time"`       //服务运行时间,纳秒
	GoroutineNum int     `json:"goroutine_num"`  //goroutine数量
	CpuNum       int     `json:"cpu_num"`        //cpu核数
	CpuUser      float64 `json:"cpu_user"`       //cpu用户态比率
	CpuFree      float64 `json:"cpu_free"`       //cpu空闲比率
	DiskUsed     uint64  `json:"disk_used"`      //已用磁盘空间,字节数
	DiskFree     uint64  `json:"disk_free"`      //可用磁盘空间,字节数
	DiskTotal    uint64  `json:"disk_total"`     //总磁盘空间,字节数
	MemUsed      uint64  `json:"mem_used"`       //已用内存,字节数
	MemSys       uint64  `json:"mem_sys"`        //系统内存占用量,字节数
	MemFree      uint64  `json:"mem_free"`       //剩余内存,字节数
	MemTotal     uint64  `json:"mem_total"`      //总内存,字节数
	AllocGolang  uint64  `json:"alloc_golang"`   //golang内存使用量,字节数
	AllocTotal   uint64  `json:"alloc_total"`    //总分配的内存,字节数
	Lookups      uint64  `json:"lookups"`        //指针查找次数
	Mallocs      uint64  `json:"mallocs"`        //内存分配次数
	Frees        uint64  `json:"frees"`          //内存释放次数
	LastGCTime   uint64  `json:"last_gc_time"`   //上次GC时间,纳秒
	NextGC       uint64  `json:"next_gc"`        //下次GC内存回收量,字节数
	PauseTotalNs uint64  `json:"pause_total_ns"` //GC暂停时间总量,纳秒
	PauseNs      uint64  `json:"pause_ns"`       //上次GC暂停时间,纳秒
}

// GetSystemInfo 获取系统运行信息.
func (toolOs *Os)SystemInfo() *RspSystem {
	//运行时信息
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)

	//CPU信息
	cpuUser, cpuIdel, cpuTotal := toolOs.CpuUsage()
	cpuUserRate := float64(cpuUser) / float64(cpuTotal)
	cpuFreeRate := float64(cpuIdel) / float64(cpuTotal)

	//磁盘空间信息
	diskUsed, diskFree, diskTotal := toolOs.DiskUsage("/")

	//内存使用信息
	memUsed, memFree, memTotal := toolOs.MemoryUsage(true)

	serverName, _ := os.Hostname()

	return &RspSystem{
		ServerName:   serverName,
		SystemOs:     runtime.GOOS,
		Runtime:      int64(time.Since(time.Now())),
		GoroutineNum: runtime.NumGoroutine(),
		CpuNum:       runtime.NumCPU(),
		CpuUser:      cpuUserRate,
		CpuFree:      cpuFreeRate,
		DiskUsed:     diskUsed,
		DiskFree:     diskFree,
		DiskTotal:    diskTotal,
		MemUsed:      memUsed,
		MemSys:       memStats.Sys,
		MemFree:      memFree,
		MemTotal:     memTotal,
		AllocGolang:  memStats.Alloc,
		AllocTotal:   memStats.TotalAlloc,
		Lookups:      memStats.Lookups,
		Mallocs:      memStats.Mallocs,
		Frees:        memStats.Frees,
		LastGCTime:   memStats.LastGC,
		NextGC:       memStats.NextGC,
		PauseTotalNs: memStats.PauseTotalNs,
		PauseNs:      memStats.PauseNs[(memStats.NumGC+255)%256],
	}
}

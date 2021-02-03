package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestSystemInfo Os/Global.go Os/SystemInfo.go Os/SystemInfo_test.go Os/CpuUsage.go Os/DiskUsage.go Os/MemoryUsage.go
func TestSystemInfo(t *testing.T) {
	var toolOs Os
	res1 := toolOs.SystemInfo()
	fmt.Println(res1)
}

// go test -v -run TestSystemInfo -bench=BenchmarkSystemInfo -count=5 Os/Global.go Os/SystemInfo.go Os/SystemInfo_test.go Os/CpuUsage.go Os/DiskUsage.go Os/MemoryUsage.go
func BenchmarkSystemInfo(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.SystemInfo()
	}
}
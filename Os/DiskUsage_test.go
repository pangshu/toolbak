package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDiskUsage Os/Global.go Os/DiskUsage.go Os/DiskUsage_test.go
func TestDiskUsage(t *testing.T) {
	var toolOs Os
	res1,res2,res3 := toolOs.DiskUsage("/")
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}

// go test -v -run TestDiskUsage -bench=BenchmarkDiskUsage -count=5 Os/Global.go Os/DiskUsage.go Os/DiskUsage_test.go
func BenchmarkDiskUsage(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_,_,_ = toolOs.DiskUsage("/")
	}
}
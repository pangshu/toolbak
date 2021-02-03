package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestCpuUsage Os/Global.go Os/CpuUsage.go Os/CpuUsage_test.go
func TestCpuUsage(t *testing.T) {
	var toolOs Os
	res1,res2,res3 := toolOs.CpuUsage()
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}

// go test -v -run TestCpuUsage -bench=BenchmarkCpuUsage -count=5 Os/Global.go Os/CpuUsage.go Os/CpuUsage_test.go
func BenchmarkCpuUsage(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_,_,_ = toolOs.CpuUsage()
	}
}
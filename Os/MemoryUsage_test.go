package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMemoryUsage Os/Global.go Os/MemoryUsage.go Os/MemoryUsage_test.go
func TestMemoryUsage(t *testing.T) {
	var toolOs Os
	res1,res2,res3 := toolOs.MemoryUsage(false)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}

// go test -v -run TestMemoryUsage -bench=BenchmarkMemoryUsage -count=5 Os/Global.go Os/MemoryUsage.go Os/MemoryUsage_test.go
func BenchmarkMemoryUsage(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_,_,_ = toolOs.MemoryUsage(false)
	}
}
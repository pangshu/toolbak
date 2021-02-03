package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestCpu Os/Global.go Os/Cpu.go Os/Cpu_test.go
func TestCpu(t *testing.T) {
	var toolOs Os
	res1 := toolOs.Cpu()
	fmt.Println(res1)
}

// go test -v -run TestCpu -bench=BenchmarkCpu -count=5 Os/Global.go Os/Cpu.go Os/Cpu_test.go
func BenchmarkCpu(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.Cpu()
	}
}
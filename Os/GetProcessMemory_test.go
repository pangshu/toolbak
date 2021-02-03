package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetProcessMemory Os/Global.go Os/GetProcessMemory.go Os/GetProcessMemory_test.go
func TestGetProcessMemory(t *testing.T) {
	var toolOs Os
	res1 := toolOs.GetProcessMemory()
	fmt.Println(res1)
}

// go test -v -run TestGetProcessMemory -bench=BenchmarkGetProcessMemory -count=5 Os/Global.go Os/GetProcessMemory.go Os/GetProcessMemory_test.go
func BenchmarkGetProcessMemory(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.GetProcessMemory()
	}
}
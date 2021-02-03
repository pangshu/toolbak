package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsWindows Os/Global.go Os/IsWindows.go Os/IsWindows_test.go
func TestIsWindows(t *testing.T) {
	var toolOs Os
	res1 := toolOs.IsWindows()
	fmt.Println(res1)
}

// go test -v -run TestIsWindows -bench=BenchmarkIsWindows -count=5 Os/Global.go Os/IsWindows.go Os/IsWindows_test.go
func BenchmarkIsWindows(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.IsWindows()
	}
}
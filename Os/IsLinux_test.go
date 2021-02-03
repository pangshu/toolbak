package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsLinux Os/Global.go Os/IsLinux.go Os/IsLinux_test.go
func TestIsLinux(t *testing.T) {
	var toolOs Os
	res1 := toolOs.IsLinux()
	fmt.Println(res1)
}

// go test -v -run TestIsLinux -bench=BenchmarkIsLinux -count=5 Os/Global.go Os/IsLinux.go Os/IsLinux_test.go
func BenchmarkIsLinux(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.IsLinux()
	}
}
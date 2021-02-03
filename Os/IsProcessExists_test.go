package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsProcessExists Os/Global.go Os/IsProcessExists.go Os/IsProcessExists_test.go
func TestIsProcessExists(t *testing.T) {
	var toolOs Os
	res1 := toolOs.IsProcessExists(15900)
	fmt.Println(res1)
}

// go test -v -run TestIsProcessExists -bench=BenchmarkIsProcessExists -count=5 Os/Global.go Os/IsProcessExists.go Os/IsProcessExists_test.go
func BenchmarkIsProcessExists(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.IsProcessExists(15900)
	}
}
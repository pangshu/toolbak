package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetTempDir Os/Global.go Os/GetTempDir.go Os/GetTempDir_test.go
func TestGetTempDir(t *testing.T) {
	var toolOs Os
	res1 := toolOs.GetTempDir()
	fmt.Println(res1)
}

// go test -v -run TestGetTempDir -bench=BenchmarkGetTempDir -count=5 Os/Global.go Os/GetTempDir.go Os/GetTempDir_test.go
func BenchmarkGetTempDir(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.GetTempDir()
	}
}
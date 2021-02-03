package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsMac Os/Global.go Os/IsMac.go Os/IsMac_test.go
func TestIsMac(t *testing.T) {
	var toolOs Os
	res1 := toolOs.IsMac()
	fmt.Println(res1)
}

// go test -v -run TestIsMac -bench=BenchmarkIsMac -count=5 Os/Global.go Os/IsMac.go Os/IsMac_test.go
func BenchmarkIsMac(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.IsMac()
	}
}
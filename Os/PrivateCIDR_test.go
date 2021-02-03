package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestPrivateCIDR Os/Global.go Os/PrivateCIDR.go Os/PrivateCIDR_test.go
func TestPrivateCIDR(t *testing.T) {
	var toolOs Os
	res1 := toolOs.PrivateCIDR()
	fmt.Println(res1)
}

// go test -v -run TestPrivateCIDR -bench=BenchmarkPrivateCIDR -count=5 Os/Global.go Os/PrivateCIDR.go Os/PrivateCIDR_test.go
func BenchmarkPrivateCIDR(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.PrivateCIDR()
	}
}
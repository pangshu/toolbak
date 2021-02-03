package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestOutBoundIp Os/Global.go Os/OutBoundIp.go Os/OutBoundIp_test.go
func TestOutBoundIp(t *testing.T) {
	var toolOs Os
	res1 := toolOs.OutBoundIp()
	fmt.Println(res1)
}

// go test -v -run TestOutBoundIp -bench=BenchmarkOutBoundIp -count=5 Os/Global.go Os/OutBoundIp.go Os/OutBoundIp_test.go
func BenchmarkOutBoundIp(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.OutBoundIp()
	}
}
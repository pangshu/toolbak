package Net

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsIPv4 Net/Global.go Net/IsIPv4.go Net/IsIPv4_test.go
func TestIsIPv4(t *testing.T) {
	var toolNet Net
	res1 := toolNet.IsIPv4("114.114.114.114")
	fmt.Println(res1)
}

// go test -v -run TestIsIPv4 -bench=BenchmarkIsIPv4 -count=5 Net/Global.go Net/IsIPv4.go Net/IsIPv4_test.go
func BenchmarkIsIPv4(t *testing.B) {
	t.ResetTimer()
	var toolNet Net
	for i:=0; i< t.N; i++ {
		_ = toolNet.IsIPv4("114.114.114.114")
	}
}
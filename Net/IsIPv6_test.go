package Net

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsIPv6 Net/Global.go Net/IsIPv6.go Net/IsIPv6_test.go
func TestIsIPv6(t *testing.T) {
	var toolNet Net
	res1 := toolNet.IsIPv6("2001:A304:6101:1::E0:F726:4E58")
	fmt.Println(res1)
}

// go test -v -run TestIsIPv6 -bench=BenchmarkIsIPv6 -count=5 Net/Global.go Net/IsIPv6.go Net/IsIPv6_test.go
func BenchmarkIsIPv6(t *testing.B) {
	t.ResetTimer()
	var toolNet Net
	for i:=0; i< t.N; i++ {
		_ = toolNet.IsIPv6("2001:A304:6101:1::E0:F726:4E58")
	}
}
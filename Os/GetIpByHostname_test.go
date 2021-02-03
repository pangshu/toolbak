package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetIpByHostname Os/Global.go Os/GetIpByHostname.go Os/GetIpByHostname_test.go
func TestGetIpByHostname(t *testing.T) {
	var toolOs Os
	ip := "127.0.0.1"
	res1,_ := toolOs.GetIpByHostname(ip)
	fmt.Println(res1)
}

// go test -v -run TestGetIpByHostname -bench=BenchmarkGetIpByHostname -count=5 Os/Global.go Os/GetIpByHostname.go Os/GetIpByHostname_test.go
func BenchmarkGetIpByHostname(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	ip := "127.0.0.1"
	for i:=0; i< t.N; i++ {
		_,_ = toolOs.GetIpByHostname(ip)
	}
}
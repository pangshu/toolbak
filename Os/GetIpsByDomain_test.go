package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetIpsByDomain Os/Global.go Os/GetIpsByDomain.go Os/GetIpsByDomain_test.go
func TestGetIpsByDomain(t *testing.T) {
	var toolOs Os
	ip := "baidu.com"
	res1,_ := toolOs.GetIpsByDomain(ip)
	fmt.Println(res1)
}

// go test -v -run TestGetIpsByDomain -bench=BenchmarkGetIpsByDomain -count=5 Os/Global.go Os/GetIpsByDomain.go Os/GetIpsByDomain_test.go
func BenchmarkGetIpsByDomain(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	ip := "baidu.com"
	for i:=0; i< t.N; i++ {
		_,_ = toolOs.GetIpsByDomain(ip)
	}
}
package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetHostByIp Os/Global.go Os/GetHostByIp.go Os/GetHostByIp_test.go
func TestGetHostByIp(t *testing.T) {
	var toolOs Os
	ip := "127.0.0.1"
	res1,_ := toolOs.GetHostByIp(ip)
	fmt.Println(res1)
}

// go test -v -run TestGetHostByIp -bench=BenchmarkGetHostByIp -count=5 Os/Global.go Os/GetHostByIp.go Os/GetHostByIp_test.go
func BenchmarkGetHostByIp(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	ip := "127.0.0.1"
	for i:=0; i< t.N; i++ {
		_,_ = toolOs.GetHostByIp(ip)
	}
}
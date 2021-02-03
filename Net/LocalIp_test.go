package Net

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetLocalIp Net/Global.go Net/GetLocalIp.go Net/GetLocalIp_test.go
func TestGetLocalIp(t *testing.T) {
	var toolNet Net
	res1 := toolNet.GetLocalIp()
	fmt.Println(res1)
}

// go test -v -run TestGetLocalIp -bench=BenchmarkGetLocalIp -count=5 Net/Global.go Net/GetLocalIp.go Net/GetLocalIp_test.go
func BenchmarkGetLocalIp(t *testing.B) {
	t.ResetTimer()
	var toolNet Net
	for i:=0; i< t.N; i++ {
		_ = toolNet.GetLocalIp()
	}
}
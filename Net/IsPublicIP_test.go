package Net

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsPublicIP Net/Global.go Net/IsPublicIP.go Net/IsPublicIP_test.go
func TestIsPublicIP(t *testing.T) {
	var toolNet Net
	res1 := toolNet.IsPublicIP("192.168.0.1")
	fmt.Println(res1)
}

// go test -v -run TestIsPublicIP -bench=BenchmarkIsPublicIP -count=5 Net/Global.go Net/IsPublicIP.go Net/IsPublicIP_test.go
func BenchmarkIsPublicIP(t *testing.B) {
	t.ResetTimer()
	var toolNet Net
	for i:=0; i< t.N; i++ {
		_ = toolNet.IsPublicIP("192.168.0.1")
	}
}
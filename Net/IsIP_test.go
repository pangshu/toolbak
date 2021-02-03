package Net

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsIP Net/Global.go Net/IsIP.go Net/IsIP_test.go
func TestIsIP(t *testing.T) {
	var toolNet Net
	res1 := toolNet.IsIP("114.114.114.114")
	fmt.Println(res1)
}

// go test -v -run TestIsIP -bench=BenchmarkIsIP -count=5 Net/Global.go Net/IsIP.go Net/IsIP_test.go
func BenchmarkIsIP(t *testing.B) {
	t.ResetTimer()
	var toolNet Net
	for i:=0; i< t.N; i++ {
		_ = toolNet.IsIP("114.114.114.114")
	}
}
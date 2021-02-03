package Net

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIpToInt Net/Global.go Net/IpToInt.go Net/IpToInt_test.go
func TestIpToInt(t *testing.T) {
	var toolNet Net
	res1 := toolNet.IpToInt("114.114.114.114")
	fmt.Println(res1)
}

// go test -v -run TestIpToInt -bench=BenchmarkIpToInt -count=5 Net/Global.go Net/IpToInt.go Net/IpToInt_test.go
func BenchmarkIpToInt(t *testing.B) {
	t.ResetTimer()
	var toolNet Net
	for i:=0; i< t.N; i++ {
		_ = toolNet.IpToInt("114.114.114.114")
	}
}
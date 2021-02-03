package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIpToLong Convert/Global.go Convert/IpToLong.go Convert/IpToLong_test.go
func TestIpToLong(t *testing.T) {
	var toolConvert Convert
	str := "192.168.1.1"
	res1 := toolConvert.IpToLong(str)
	fmt.Println(res1)
}

// go test -v -run TestIpToLong -bench=BenchmarkIpToLong -count=5 Convert/Global.go Convert/IpToLong.go Convert/IpToLong_test.go
func BenchmarkIpToLong(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "192.168.1.1"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IpToLong(str)
	}
}
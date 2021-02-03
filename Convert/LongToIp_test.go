package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestLongToIp Convert/Global.go Convert/LongToIp.go Convert/LongToIp_test.go
func TestLongToIp(t *testing.T) {
	var toolConvert Convert
	str := uint32(3232235777)
	res1 := toolConvert.LongToIp(str)
	fmt.Println(res1)
}

// go test -v -run TestLongToIp -bench=BenchmarkLongToIp -count=5 Convert/Global.go Convert/LongToIp.go Convert/LongToIp_test.go
func BenchmarkLongToIp(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := uint32(3232235777)
	for i:=0; i< t.N; i++ {
		_ = toolConvert.LongToIp(str)
	}
}
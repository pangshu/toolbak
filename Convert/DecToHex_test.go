package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDecToHex Convert/Global.go Convert/DecToHex.go Convert/DecToHex_test.go
func TestDecToHex(t *testing.T) {
	var toolConvert Convert
	var str int64 = 11
	res1 := toolConvert.DecToHex(str)
	fmt.Println(res1)
}

// go test -v -run TestDecToHex -bench=BenchmarkDecToHex -count=5 Convert/Global.go Convert/DecToHex.go Convert/DecToHex_test.go
func BenchmarkDecToHex(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	var str int64 = 11
	for i:=0; i< t.N; i++ {
		_ = toolConvert.DecToHex(str)
	}
}
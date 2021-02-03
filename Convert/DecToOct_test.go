package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDecToOct Convert/Global.go Convert/DecToOct.go Convert/DecToOct_test.go
func TestDecToOct(t *testing.T) {
	var toolConvert Convert
	var str int64 = 11
	res1 := toolConvert.DecToOct(str)
	fmt.Println(res1)
}

// go test -v -run TestDecToOct -bench=BenchmarkDecToOct -count=5 Convert/Global.go Convert/DecToOct.go Convert/DecToOct_test.go
func BenchmarkDecToOct(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	var str int64 = 11
	for i:=0; i< t.N; i++ {
		_ = toolConvert.DecToOct(str)
	}
}
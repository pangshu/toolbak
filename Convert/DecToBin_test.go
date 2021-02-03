package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDecToBin Convert/Global.go Convert/DecToBin.go Convert/DecToBin_test.go
func TestDecToBin(t *testing.T) {
	var toolConvert Convert
	var str int64 = 8
	res1 := toolConvert.DecToBin(str)
	fmt.Println(res1)
}

// go test -v -run TestDecToBin -bench=BenchmarkDecToBin -count=5 Convert/Global.go Convert/DecToBin.go Convert/DecToBin_test.go
func BenchmarkDecToBin(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	var str int64 = 8
	for i:=0; i< t.N; i++ {
		_ = toolConvert.DecToBin(str)
	}
}
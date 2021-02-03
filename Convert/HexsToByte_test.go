package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestHexsToByte Convert/Global.go Convert/HexsToByte.go Convert/HexsToByte_test.go
func TestHexsToByte(t *testing.T) {
	var toolConvert Convert
	str := []byte("68656c6c6f")
	res1 := toolConvert.HexsToByte(str)
	fmt.Println(res1)
}

// go test -v -run TestHexsToByte -bench=BenchmarkHexsToByte -count=5 Convert/Global.go Convert/HexsToByte.go Convert/HexsToByte_test.go
func BenchmarkHexsToByte(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := []byte("68656c6c6f")
	for i:=0; i< t.N; i++ {
		_ = toolConvert.HexsToByte(str)
	}
}
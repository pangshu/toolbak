package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestHexToByte Convert/Global.go Convert/HexToByte.go Convert/HexToByte_test.go
func TestHexToByte(t *testing.T) {
	var toolConvert Convert
	str := "68656c6c6f"
	res1 := toolConvert.HexToByte(str)
	fmt.Println(string(res1))
}

// go test -v -run TestHexToByte -bench=BenchmarkHexToByte -count=5 Convert/Global.go Convert/HexToByte.go Convert/HexToByte_test.go
func BenchmarkHexToByte(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "68656c6c6f"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.HexToByte(str)
	}
}
package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestHexToDec Convert/Global.go Convert/HexToDec.go Convert/HexToDec_test.go
func TestHexToDec(t *testing.T) {
	var toolConvert Convert
	str := "68656c6c6f"
	res1,_ := toolConvert.HexToDec(str)
	fmt.Println(res1)
}

// go test -v -run TestHexToDec -bench=BenchmarkHexToDec -count=5 Convert/Global.go Convert/HexToDec.go Convert/HexToDec_test.go
func BenchmarkHexToDec(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "68656c6c6f"
	for i:=0; i< t.N; i++ {
		_,_ = toolConvert.HexToDec(str)
	}
}
package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestHexToBin Convert/Global.go Convert/HexToBin.go Convert/HexToBin_test.go
func TestHexToBin(t *testing.T) {
	var toolConvert Convert
	str := "a"
	res1,_ := toolConvert.HexToBin(str)
	fmt.Println(res1)
}

// go test -v -run TestHexToBin -bench=BenchmarkHexToBin -count=5 Convert/Global.go Convert/HexToBin.go Convert/HexToBin_test.go
func BenchmarkHexToBin(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "a"
	for i:=0; i< t.N; i++ {
		_,_ = toolConvert.HexToBin(str)
	}
}
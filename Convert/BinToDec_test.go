package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestBinToDec Convert/Global.go Convert/BinToDec.go Convert/BinToDec_test.go
func TestBinToDec(t *testing.T) {
	var toolConvert Convert
	str := "11011101"
	res1,err := toolConvert.BinToDec(str)
	fmt.Println(res1)
	fmt.Println(err)
}

// go test -v -run TestBinToDec -bench=BenchmarkBinToDec -count=5 Convert/Global.go Convert/BinToDec.go Convert/BinToDec_test.go
func BenchmarkBinToDec(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "11011101"
	for i:=0; i< t.N; i++ {
		_,_ = toolConvert.BinToDec(str)
	}
}
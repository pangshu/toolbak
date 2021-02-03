package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestBinToHex Convert/Global.go Convert/BinToHex.go Convert/BinToHex_test.go
func TestBinToHex(t *testing.T) {
	var toolConvert Convert
	str := "11011101"
	res1,err := toolConvert.BinToHex(str)
	fmt.Println(res1)
	fmt.Println(err)
}

// go test -v -run TestBinToHex -bench=BenchmarkBinToHex -count=5 Convert/Global.go Convert/BinToHex.go Convert/BinToHex_test.go
func BenchmarkBinToHex(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "11011101"
	for i:=0; i< t.N; i++ {
		_,_ = toolConvert.BinToHex(str)
	}
}
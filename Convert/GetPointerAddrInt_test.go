package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetPointerAddrInt Convert/Global.go Convert/GetPointerAddrInt.go Convert/GetPointerAddrInt_test.go Convert/HexToDec.go
func TestGetPointerAddrInt(t *testing.T) {
	var toolConvert Convert
	str := []byte("hello world!")
	res1 := toolConvert.GetPointerAddrInt(str)
	fmt.Println(res1)
}

// go test -v -run TestGetPointerAddrInt -bench=BenchmarkGetPointerAddrInt -count=5 Convert/Global.go Convert/GetPointerAddrInt.go Convert/GetPointerAddrInt_test.go Convert/HexToDec.go
func BenchmarkGetPointerAddrInt(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := []byte("hello world!")
	for i:=0; i< t.N; i++ {
		_ = toolConvert.GetPointerAddrInt(str)
	}
}
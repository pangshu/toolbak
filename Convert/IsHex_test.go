package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsHex Convert/Global.go Convert/IsHex.go Convert/IsHex_test.go Convert/HexToDec.go
func TestIsHex(t *testing.T) {
	var toolConvert Convert
	str := "123456789ABCDEF"
	res1 := toolConvert.IsHex(str)
	fmt.Println(res1)
}

// go test -v -run TestIsHex -bench=BenchmarkIsHex -count=5 Convert/Global.go Convert/IsHex.go Convert/IsHex_test.go Convert/HexToDec.go
func BenchmarkIsHex(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456789ABCDEF"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IsHex(str)
	}
}
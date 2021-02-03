package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestByteToHex Convert/Global.go Convert/ByteToHex.go Convert/ByteToHex_test.go
func TestByteToHex(t *testing.T) {
	var toolConvert Convert
	str := []byte("hello word!")
	res1 := toolConvert.ByteToHex(str)
	fmt.Println(res1)
}

// go test -v -run TestByteToHex -bench=BenchmarkByteToHex -count=5 Convert/Global.go Convert/ByteToHex.go Convert/ByteToHex_test.go
func BenchmarkByteToHex(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := []byte("hello word!")
	for i:=0; i< t.N; i++ {
		_ = toolConvert.ByteToHex(str)
	}
}
package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestByteToBase64 Convert/Global.go Convert/ByteToBase64.go Convert/ByteToBase64_test.go
func TestByteToBase64(t *testing.T) {
	var toolConvert Convert
	str := []byte("hello world!")
	res1 := toolConvert.ByteToBase64(str)
	fmt.Println(res1)
}

// go test -v -run TestByteToBase64 -bench=BenchmarkByteToBase64 -count=5 Convert/Global.go Convert/ByteToBase64.go Convert/ByteToBase64_test.go
func BenchmarkByteToBase64(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := []byte("hello world!")
	for i:=0; i< t.N; i++ {
		_ = toolConvert.ByteToBase64(str)
	}
}
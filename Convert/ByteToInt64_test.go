package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestByteToInt64 Convert/Global.go Convert/ByteToInt64.go Convert/ByteToInt64_test.go
func TestByteToInt64(t *testing.T) {
	var toolConvert Convert
	str := []byte("hello word!")
	res1 := toolConvert.ByteToInt64(str)
	fmt.Println(res1)
}

// go test -v -run TestByteToInt64 -bench=BenchmarkByteToInt64 -count=5 Convert/Global.go Convert/ByteToInt64.go Convert/ByteToInt64_test.go
func BenchmarkByteToInt64(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := []byte("hello word!")
	for i:=0; i< t.N; i++ {
		_ = toolConvert.ByteToInt64(str)
	}
}
package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestByteToFloat64 Convert/Global.go Convert/ByteToFloat64.go Convert/ByteToFloat64_test.go
func TestByteToFloat64(t *testing.T) {
	var toolConvert Convert
	str := []byte{205, 204, 204, 204, 204, 28, 200, 64}
	res1 := toolConvert.ByteToFloat64(str)
	fmt.Println(res1)
}

// go test -v -run TestByteToFloat64 -bench=BenchmarkByteToFloat64 -count=5 Convert/Global.go Convert/ByteToFloat64.go Convert/ByteToFloat64_test.go
func BenchmarkByteToFloat64(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := []byte{205, 204, 204, 204, 204, 28, 200, 64}
	for i:=0; i< t.N; i++ {
		_ = toolConvert.ByteToFloat64(str)
	}
}
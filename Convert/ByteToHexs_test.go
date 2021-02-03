package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestByteToHexs Convert/Global.go Convert/ByteToHexs.go Convert/ByteToHexs_test.go
func TestByteToHexs(t *testing.T) {
	var toolConvert Convert
	str := []byte("hello word!")
	res1 := toolConvert.ByteToHexs(str)
	fmt.Println(res1)
}

// go test -v -run TestByteToHexs -bench=BenchmarkByteToHexs -count=5 Convert/Global.go Convert/ByteToHexs.go Convert/ByteToHexs_test.go
func BenchmarkByteToHexs(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := []byte("hello word!")
	for i:=0; i< t.N; i++ {
		_ = toolConvert.ByteToHexs(str)
	}
}
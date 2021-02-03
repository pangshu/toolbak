package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestFromBytes Convert/Global.go Convert/FromBytes.go Convert/FromBytes_test.go
func TestFromBytes(t *testing.T) {
	var toolConvert Convert
	str := []byte("hello world!")
	res1 := toolConvert.FromBytes(str)
	fmt.Println(res1)
}

// go test -v -run TestFromBytes -bench=BenchmarkFromBytes -count=5 Convert/Global.go Convert/FromBytes.go Convert/FromBytes_test.go
func BenchmarkFromBytes(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := []byte("hello world!")
	for i:=0; i< t.N; i++ {
		_ = toolConvert.FromBytes(str)
	}
}
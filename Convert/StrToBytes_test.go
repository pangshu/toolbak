package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToBytes Convert/Global.go Convert/StrToBytes.go Convert/StrToBytes_test.go
func TestStrToBytes(t *testing.T) {
	var toolConvert Convert
	str := "hello world!"
	res1 := toolConvert.StrToBytes(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToBytes -bench=BenchmarkStrToBytes -count=5 Convert/Global.go Convert/StrToBytes.go Convert/StrToBytes_test.go
func BenchmarkStrToBytes(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "hello world!"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToBytes(str)
	}
}
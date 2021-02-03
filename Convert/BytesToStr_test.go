package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestBytesToStr Convert/Global.go Convert/BytesToStr.go Convert/BytesToStr_test.go
func TestBytesToStr(t *testing.T) {
	var toolConvert Convert
	str := []byte("hello world!")
	res1 := toolConvert.BytesToStr(str)
	fmt.Println(res1)
}

// go test -v -run TestBytesToStr -bench=BenchmarkBytesToStr -count=5 Convert/Global.go Convert/BytesToStr.go Convert/BytesToStr_test.go
func BenchmarkBytesToStr(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := []byte("hello world!")
	for i:=0; i< t.N; i++ {
		_ = toolConvert.BytesToStr(str)
	}
}
package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGettype Convert/Global.go Convert/Gettype.go Convert/Gettype_test.go
func TestGettype(t *testing.T) {
	var toolConvert Convert
	str := []byte("hello world!")
	res1 := toolConvert.Gettype(str)
	fmt.Println(res1)
}

// go test -v -run TestGettype -bench=BenchmarkGettype -count=5 Convert/Global.go Convert/Gettype.go Convert/Gettype_test.go
func BenchmarkGettype(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := []byte("hello world!")
	for i:=0; i< t.N; i++ {
		_ = toolConvert.Gettype(str)
	}
}
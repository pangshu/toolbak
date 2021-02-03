package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsBinary Convert/Global.go Convert/IsBinary.go Convert/IsBinary_test.go
func TestIsBinary(t *testing.T) {
	var toolConvert Convert
	str := "110110110"
	res1 := toolConvert.IsBinary(str)
	fmt.Println(res1)
}

// go test -v -run TestIsBinary -bench=BenchmarkIsBinary -count=5 Convert/Global.go Convert/IsBinary.go Convert/IsBinary_test.go
func BenchmarkIsBinary(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "110110110"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IsBinary(str)
	}
}
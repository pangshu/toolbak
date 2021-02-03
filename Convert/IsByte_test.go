package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsByte Convert/Global.go Convert/IsByte.go Convert/IsByte_test.go Convert/Gettype.go
func TestIsByte(t *testing.T) {
	var toolConvert Convert
	res1 := toolConvert.IsByte([]byte("hello"))
	fmt.Println(res1)
}

// go test -v -run TestIsByte -bench=BenchmarkIsByte -count=5 Convert/Global.go Convert/IsByte.go Convert/IsByte_test.go Convert/Gettype.go
func BenchmarkIsByte(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IsByte([]byte("hello"))
	}
}
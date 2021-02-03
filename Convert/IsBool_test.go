package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsBool Convert/Global.go Convert/IsBool.go Convert/IsBool_test.go
func TestIsBool(t *testing.T) {
	var toolConvert Convert
	res1 := toolConvert.IsBool(true)
	fmt.Println(res1)
}

// go test -v -run TestIsBool -bench=BenchmarkIsBool -count=5 Convert/Global.go Convert/IsBool.go Convert/IsBool_test.go
func BenchmarkIsBool(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IsBool(true)
	}
}
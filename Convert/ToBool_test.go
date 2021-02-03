package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestToBool Convert/Global.go Convert/ToBool.go Convert/ToBool_test.go
func TestToBool(t *testing.T) {
	var toolConvert Convert
	str := "1"
	res1 := toolConvert.ToBool(str)
	fmt.Println(res1)
}

// go test -v -run TestToBool -bench=BenchmarkToBool -count=5 Convert/Global.go Convert/ToBool.go Convert/ToBool_test.go
func BenchmarkToBool(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "1"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.ToBool(str)
	}
}
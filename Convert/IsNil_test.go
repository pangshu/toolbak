package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsNil Convert/Global.go Convert/IsNil.go Convert/IsNil_test.go
func TestIsNil(t *testing.T) {
	var toolConvert Convert
	str := ""
	res1 := toolConvert.IsNil(str)
	fmt.Println(res1)
}

// go test -v -run TestIsNil -bench=BenchmarkIsNil -count=5 Convert/Global.go Convert/IsNil.go Convert/IsNil_test.go
func BenchmarkIsNil(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := ""
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IsNil(str)
	}
}
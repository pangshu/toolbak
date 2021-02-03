package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestToInterface Convert/Global.go Convert/ToInterface.go Convert/ToInterface_test.go
func TestToInterface(t *testing.T) {
	var toolConvert Convert
	str := []string{"123","234"}
	res1 := toolConvert.ToInterface(str)
	fmt.Println(res1)
}

// go test -v -run TestToInterface -bench=BenchmarkToInterface -count=5 Convert/Global.go Convert/ToInterface.go Convert/ToInterface_test.go
func BenchmarkToInterface(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := []string{"123","234"}
	for i:=0; i< t.N; i++ {
		_ = toolConvert.ToInterface(str)
	}
}
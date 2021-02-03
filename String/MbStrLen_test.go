package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMbStrLen String/Global.go String/MbStrLen.go String/MbStrLen_test.go
func TestMbStrLen(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.MbStrLen("中文aa bb cc AaBbCcDd")
	fmt.Println(res1)
}

// go test -v -run TestMbStrLen -bench=BenchmarkMbStrLen -count=5 String/Global.go String/MbStrLen.go String/MbStrLen_test.go
func BenchmarkMbStrLen(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.MbStrLen("中文aa bb cc AaBbCcDd")
	}
}
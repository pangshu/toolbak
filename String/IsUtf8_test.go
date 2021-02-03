package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsUtf8 String/Global.go String/IsUtf8.go String/IsUtf8_test.go
func TestIsUtf8(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.IsUtf8("中文aa bb cc AaBbCcDd")
	fmt.Println(res1)
}

// go test -v -run TestIsUtf8 -bench=BenchmarkIsUtf8 -count=5 String/Global.go String/IsUtf8.go String/IsUtf8_test.go
func BenchmarkIsUtf8(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.IsUtf8("中文aa bb cc AaBbCcDd")
	}
}
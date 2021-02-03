package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDelSpace String/Global.go String/DelSpace.go String/DelSpace_test.go
func TestDelSpace(t *testing.T) {
	var toolbaktring String
	str := "Hello 你好, World 世界！"
	res1 := toolbaktring.DelSpace(str, true)
	fmt.Println(res1)
}

// go test -v -run TestDelSpace -bench=BenchmarkDelSpace -count=5 String/Global.go String/DelSpace.go String/DelSpace_test.go
func BenchmarkDelSpace(t *testing.B) {
	t.ResetTimer()
	str := "Hello 你好, World 世界！"
	var toolbaktring String
	for i:=0; i< t.N; i++ {
		_ = toolbaktring.DelSpace(str, true)
	}
}
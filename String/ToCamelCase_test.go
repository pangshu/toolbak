package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestToCamelCase String/Global.go String/ToCamelCase.go String/ToCamelCase_test.go
func TestToCamelCase(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.ToCamelCase("aa bb cc AaBbCcDd")
	fmt.Println(res1)
}

// go test -v -run TestToCamelCase -bench=BenchmarkToCamelCase -count=5 String/Global.go String/ToCamelCase.go String/ToCamelCase_test.go
func BenchmarkToCamelCase(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.ToCamelCase("aa bb cc AaBbCcDd")
	}
}
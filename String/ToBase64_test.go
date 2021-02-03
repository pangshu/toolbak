package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestToBase64 String/Global.go String/ToBase64.go String/ToBase64_test.go
func TestToBase64(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.ToBase64("aa bb cc AaBbCcDd")
	fmt.Println(res1)
}

// go test -v -run TestToBase64 -bench=BenchmarkToBase64 -count=5 String/Global.go String/ToBase64.go String/ToBase64_test.go
func BenchmarkToBase64(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.ToBase64("aa bb cc AaBbCcDd")
	}
}
package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestSubstr String/Global.go String/Substr.go String/Substr_test.go
func TestSubstr(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.Substr("中文aa bb cc AaBbCcDd", 0, 6)
	fmt.Println(res1)
}

// go test -v -run TestSubstr -bench=BenchmarkSubstr -count=5 String/Global.go String/Substr.go String/Substr_test.go
func BenchmarkSubstr(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.Substr("中文aa bb cc AaBbCcDd", 0, 6)
	}
}
package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMbSubstr String/Global.go String/MbSubstr.go String/MbSubstr_test.go
func TestMbSubstr(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.MbSubstr("中文aa bb cc AaBbCcDd", 0, 6)
	fmt.Println(res1)
}

// go test -v -run TestMbSubstr -bench=BenchmarkMbSubstr -count=5 String/Global.go String/MbSubstr.go String/MbSubstr_test.go
func BenchmarkMbSubstr(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.MbSubstr("中文aa bb cc AaBbCcDd", 0, 6)
	}
}
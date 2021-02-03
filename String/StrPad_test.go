package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrPad String/Global.go String/StrPad.go String/StrPad_test.go
func TestStrPad(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.StrPad("aa bb cc AaBbCcDd", "------", 50, "ALL")
	fmt.Println(res1)
}

// go test -v -run TestStrPad -bench=BenchmarkStrPad -count=5 String/Global.go String/StrPad.go String/StrPad_test.go
func BenchmarkStrPad(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.StrPad("aa bb cc AaBbCcDd", "------", 50, "ALL")
	}
}
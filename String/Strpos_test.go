package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrpos String/Global.go String/Strpos.go String/Strpos_test.go
func TestStrpos(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.Strpos("aa bb cc AaBbCcDd", "AaB", 0)
	fmt.Println(res1)
}

// go test -v -run TestStrpos -bench=BenchmarkStrpos -count=5 String/Global.go String/Strpos.go String/Strpos_test.go
func BenchmarkStrpos(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.Strpos("aa bb cc AaBbCcDd", "AaB", 1)
	}
}
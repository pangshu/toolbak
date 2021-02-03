package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestTrimLeft String/Global.go String/TrimLeft.go String/TrimLeft_test.go
func TestTrimLeft(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.TrimLeft("  半角空格　")
	res2 := toolbaktring.TrimLeft("　全角空格")
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(len("  半角空格　"), " >>>> ", len(res1))
}

// go test -v -run TestTrimLeft -bench=BenchmarkTrimLeft -count=5 String/Global.go String/TrimLeft.go String/TrimLeft_test.go
func BenchmarkTrimLeft(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.TrimLeft("  半角空格")
		_ = toolbaktring.TrimLeft("　全角空格")
	}
}
package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestTrimRight String/Global.go String/TrimRight.go String/TrimRight_test.go
func TestTrimRight(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.TrimRight("  半角空格　")
	res2 := toolbaktring.TrimRight("　全角空格")
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(len("  半角空格　"), " >>>> ", len(res1))
}

// go test -v -run TestTrimRight -bench=BenchmarkTrimRight -count=5 String/Global.go String/TrimRight.go String/TrimRight_test.go
func BenchmarkTrimRight(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.TrimRight("  半角空格")
		_ = toolbaktring.TrimRight("　全角空格")
	}
}
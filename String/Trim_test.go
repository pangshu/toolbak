package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestTrim String/Trim_test.go String/Global.go String/Trim.go
func TestTrim(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.Trim("  半角空格　")
	res2 := toolbaktring.Trim("　全角空格")
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(len("  半角空格　"), " >>>> ", len(res1))
}

// go test -v -run TestTrim -bench=BenchmarkTrim -count=5 String/Trim_test.go String/Global.go String/Trim.go
func BenchmarkTrim(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.Trim("  半角空格")
		_ = toolbaktring.Trim("　全角空格")
	}
}
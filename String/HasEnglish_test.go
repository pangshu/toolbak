package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestHasEnglish String/Global.go String/HasEnglish.go String/HasEnglish_test.go String/HasLetter.go
func TestHasEnglish(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.HasEnglish("aaa bbb 一 ddd")
	fmt.Println(res1)
}

// go test -v -run TestHasEnglish -bench=BenchmarkHasEnglish -count=5 String/Global.go String/HasEnglish.go String/HasEnglish_test.go String/HasLetter.go
func BenchmarkHasEnglish(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.HasEnglish("aaa bbb 一 ddd")
	}
}
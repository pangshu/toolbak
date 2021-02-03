package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestHasLetter String/Global.go String/HasLetter.go String/HasLetter_test.go
func TestHasLetter(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.HasLetter("aaa bbb 一 ddd")
	fmt.Println(res1)
}

// go test -v -run TestHasLetter -bench=BenchmarkHasLetter -count=5 String/Global.go String/HasLetter.go String/HasLetter_test.go
func BenchmarkHasLetter(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.HasLetter("aaa bbb 一 ddd")
	}
}
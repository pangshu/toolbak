package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestToHtml String/Global.go String/ToHtml.go String/ToHtml_test.go
func TestToHtml(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.ToHtml("hi, what's your name? 您叫什么名字？")
	fmt.Println(res1)
}

// go test -v -run TestToHtml -bench=BenchmarkToHtml -count=5 String/Global.go String/ToHtml.go String/ToHtml_test.go
func BenchmarkToHtml(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.ToHtml("hi, what's your name? 您叫什么名字？")
	}
}
package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestSimilarText String/Global.go String/SimilarText.go String/SimilarText_test.go
func TestSimilarText(t *testing.T) {
	var toolbaktring String
	var percent float64
	res1 := toolbaktring.SimilarText("你好吗，饭吃了吗？", "你好吗，饭吃了没？", &percent)
	fmt.Println(res1)
	fmt.Println(percent)
}

// go test -v -run TestSimilarText -bench=BenchmarkSimilarText -count=5 String/Global.go String/SimilarText.go String/SimilarText_test.go
func BenchmarkSimilarText(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		var percent float64
		_ = toolbaktring.SimilarText("你好吗，饭吃了吗？", "你好吗，饭吃了没？", &percent)
	}
}
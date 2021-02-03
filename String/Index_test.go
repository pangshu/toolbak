package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIndex String/Global.go String/Index.go String/Index_test.go
func TestIndex(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.Index("Who's Bill Gates?", "bill", true)
	fmt.Println(res1)
}

// go test -v -run TestIndex -bench=BenchmarkIndex -count=5 String/Global.go String/Index.go String/Index_test.go
func BenchmarkIndex(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.Index("Who's Bill Gates?","Who")
	}
}
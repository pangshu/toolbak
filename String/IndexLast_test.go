package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIndexLast String/Global.go String/IndexLast.go String/IndexLast_test.go
func TestIndexLast(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.IndexLast("Who's Bill Gates? Bill Gates?", "bill", true)
	fmt.Println(res1)
}

// go test -v -run TestIndexLast -bench=BenchmarkIndexLast -count=5 String/Global.go String/IndexLast.go String/IndexLast_test.go
func BenchmarkIndexLast(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.IndexLast("Who's Bill Gates? Bill Gates?","Who")
	}
}
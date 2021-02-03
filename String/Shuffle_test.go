package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestShuffle String/Global.go String/Shuffle.go String/Shuffle_test.go
func TestShuffle(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.Shuffle("Who's Bill Gates?")
	fmt.Println(res1)
}

// go test -v -run TestShuffle -bench=BenchmarkShuffle -count=5 String/Global.go String/Shuffle.go String/Shuffle_test.go
func BenchmarkShuffle(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.Shuffle("Who's Bill Gates?")
	}
}
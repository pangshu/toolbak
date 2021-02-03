package Rand

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIntNoRepeat Rand/Global.go Rand/IntNoRepeat.go Rand/IntNoRepeat_test.go
func TestIntNoRepeat(t *testing.T) {
	var toolRand Rand
	res1 := toolRand.IntNoRepeat(1,100, 10)
	fmt.Println(res1)
}

// go test -v -run TestIntNoRepeat -bench=BenchmarkIntNoRepeat -count=5 Rand/Global.go Rand/IntNoRepeat.go Rand/IntNoRepeat_test.go
func BenchmarkIntNoRepeat(t *testing.B) {
	t.ResetTimer()
	var toolRand Rand
	for i:=0; i< t.N; i++ {
		_ = toolRand.IntNoRepeat(1,100, 10)
	}
}
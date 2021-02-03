package Rand

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIntArray Rand/Global.go Rand/IntArray.go Rand/IntArray_test.go
func TestIntArray(t *testing.T) {
	var toolRand Rand
	res1 := toolRand.IntArray(1,100, 10)
	fmt.Println(res1)
}

// go test -v -run TestIntArray -bench=BenchmarkIntArray -count=5 Rand/Global.go Rand/IntArray.go Rand/IntArray_test.go
func BenchmarkIntArray(t *testing.B) {
	t.ResetTimer()
	var toolRand Rand
	for i:=0; i< t.N; i++ {
		_ = toolRand.IntArray(1,100, 10)
	}
}
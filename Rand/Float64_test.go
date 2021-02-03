package Rand

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestFloat64 Rand/Global.go Rand/Float64.go Rand/Float64_test.go
func TestFloat64(t *testing.T) {
	var toolRand Rand
	res1 := toolRand.Float64(1,10)
	fmt.Println(res1)
}

// go test -v -run TestFloat64 -bench=BenchmarkFloat64 -count=5 Rand/Global.go Rand/Float64.go Rand/Float64_test.go
func BenchmarkFloat64(t *testing.B) {
	t.ResetTimer()
	var toolRand Rand
	for i:=0; i< t.N; i++ {
		_ = toolRand.Float64(1,10)
	}
}
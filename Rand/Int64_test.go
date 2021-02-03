package Rand

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestInt64 Rand/Global.go Rand/Int64.go Rand/Int64_test.go
func TestInt64(t *testing.T) {
	var toolRand Rand
	res1 := toolRand.Int64(1,10)
	fmt.Println(res1)
}

// go test -v -run TestInt64 -bench=BenchmarkInt64 -count=5 Rand/Global.go Rand/Int64.go Rand/Int64_test.go
func BenchmarkInt64(t *testing.B) {
	t.ResetTimer()
	var toolRand Rand
	for i:=0; i< t.N; i++ {
		_ = toolRand.Int64(1,10)
	}
}
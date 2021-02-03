package Rand

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestString Rand/Global.go Rand/String.go Rand/String_test.go
func TestString(t *testing.T) {
	var toolRand Rand
	res1 := toolRand.String(10)
	fmt.Println(res1)
}

// go test -v -run TestString -bench=BenchmarkString -count=5 Rand/Global.go Rand/String.go Rand/String_test.go
func BenchmarkString(t *testing.B) {
	t.ResetTimer()
	var toolRand Rand
	for i:=0; i< t.N; i++ {
		_ = toolRand.String(10)
	}
}
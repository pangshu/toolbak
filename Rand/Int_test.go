package Rand

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestInt Rand/Global.go Rand/Int.go Rand/Int_test.go
func TestInt(t *testing.T) {
	var toolRand Rand
	res1 := toolRand.Int(1,10)
	fmt.Println(res1)
}

// go test -v -run TestInt -bench=BenchmarkInt -count=5 Rand/Global.go Rand/Int.go Rand/Int_test.go
func BenchmarkInt(t *testing.B) {
	t.ResetTimer()
	var toolRand Rand
	for i:=0; i< t.N; i++ {
		_ = toolRand.Int(1,10)
	}
}
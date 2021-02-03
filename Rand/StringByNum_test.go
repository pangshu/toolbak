package Rand

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStringByNum Rand/Global.go Rand/StringByNum.go Rand/StringByNum_test.go
func TestStringByNum(t *testing.T) {
	var toolRand Rand
	res1 := toolRand.StringByNum(10)
	fmt.Println(res1)
}

// go test -v -run TestStringByNum -bench=BenchmarkStringByNum -count=5 Rand/Global.go Rand/StringByNum.go Rand/StringByNum_test.go
func BenchmarkStringByNum(t *testing.B) {
	t.ResetTimer()
	var toolRand Rand
	for i:=0; i< t.N; i++ {
		_ = toolRand.StringByNum(10)
	}
}
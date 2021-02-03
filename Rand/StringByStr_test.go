package Rand

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStringByStr Rand/Global.go Rand/StringByStr.go Rand/StringByStr_test.go
func TestStringByStr(t *testing.T) {
	var toolRand Rand
	res1 := toolRand.StringByStr(10)
	fmt.Println(res1)
}

// go test -v -run TestStringByStr -bench=BenchmarkStringByStr -count=5 Rand/Global.go Rand/StringByStr.go Rand/StringByStr_test.go
func BenchmarkStringByStr(t *testing.B) {
	t.ResetTimer()
	var toolRand Rand
	for i:=0; i< t.N; i++ {
		_ = toolRand.StringByStr(10)
	}
}
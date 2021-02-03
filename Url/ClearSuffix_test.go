package Url

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestClearSuffix Url/Global.go Url/ClearSuffix.go Url/ClearSuffix_test.go
func TestClearSuffix(t *testing.T) {
	var toolUrl Url
	str := "https://github.com/abc"
	arr := "com/abc"
	res1 := toolUrl.ClearSuffix(str,arr)
	fmt.Println(res1)
}

// go test -v -run TestClearSuffix -bench=BenchmarkClearSuffix -count=5 Url/Global.go Url/ClearSuffix.go Url/ClearSuffix_test.go
func BenchmarkClearSuffix(t *testing.B) {
	t.ResetTimer()
	str := "https://github.com/abc"
	arr := "com/abc"
	var toolUrl Url
	for i:=0; i< t.N; i++ {
		_ = toolUrl.ClearSuffix(str,arr)
	}
}
package Url

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestClearPrefix Url/Global.go Url/ClearPrefix.go Url/ClearPrefix_test.go
func TestClearPrefix(t *testing.T) {
	var toolUrl Url
	str := "https://github.com/abc"
	arr := "https://"
	res1 := toolUrl.ClearPrefix(str,arr)
	fmt.Println(res1)
}

// go test -v -run TestClearPrefix -bench=BenchmarkClearPrefix -count=5 Url/Global.go Url/ClearPrefix.go Url/ClearPrefix_test.go
func BenchmarkClearPrefix(t *testing.B) {
	t.ResetTimer()
	str := "https://github.com/abc"
	arr := "https://"
	var toolUrl Url
	for i:=0; i< t.N; i++ {
		_ = toolUrl.ClearPrefix(str,arr)
	}
}
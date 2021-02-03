package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestHasWhitespace String/Global.go String/HasWhitespace.go String/HasWhitespace_test.go
func TestHasWhitespace(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.HasWhitespace("aaa bbb 一 ddd&")
	fmt.Println(res1)
}

// go test -v -run TestHasWhitespace -bench=BenchmarkHasWhitespace -count=5 String/Global.go String/HasWhitespace.go String/HasWhitespace_test.go
func BenchmarkHasWhitespace(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.HasWhitespace("aaa bbb 一 ddd")
	}
}
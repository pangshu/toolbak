package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestHasFullWidth String/Global.go String/HasFullWidth.go String/HasFullWidth_test.go
func TestHasFullWidth(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.HasFullWidth("aaa bbb 一 ddd dd　dd")
	fmt.Println(res1)
}

// go test -v -run TestHasFullWidth -bench=BenchmarkHasFullWidth -count=5 String/Global.go String/HasFullWidth.go String/HasFullWidth_test.go
func BenchmarkHasFullWidth(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.HasFullWidth("aaa bbb 一 ddd dd　dd")
	}
}
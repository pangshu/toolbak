package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestHasChinese String/Global.go String/HasChinese.go String/HasChinese_test.go
func TestHasChinese(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.HasChinese("aaa bbb 一 ddd")
	fmt.Println(res1)
}

// go test -v -run TestHasChinese -bench=BenchmarkHasChinese -count=5 String/Global.go String/HasChinese.go String/HasChinese_test.go
func BenchmarkHasChinese(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.HasChinese("aaa bbb 一 ddd")
	}
}
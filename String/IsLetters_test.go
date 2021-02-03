package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsLetters String/Global.go String/IsLetters.go String/IsLetters_test.go
func TestIsLetters(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.IsLetters("aabbcc")
	fmt.Println(res1)
}

// go test -v -run TestIsLetters -bench=BenchmarkIsLetters -count=5 String/Global.go String/IsLetters.go String/IsLetters_test.go
func BenchmarkIsLetters(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.IsLetters("中文aa bb cc AaBbCcDd")
	}
}
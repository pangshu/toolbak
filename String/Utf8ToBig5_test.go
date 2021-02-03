package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestUtf8ToBig5 String/Global.go String/Utf8ToBig5.go String/Utf8ToBig5_test.go
func TestUtf8ToBig5(t *testing.T) {
	var toolbaktring String
	str := "你好，世界！"
	res1,_ := toolbaktring.Utf8ToBig5(str)
	fmt.Println(res1)
}

// go test -v -run TestUtf8ToBig5 -bench=BenchmarkUtf8ToBig5 -count=5 String/Global.go String/Utf8ToBig5.go String/Utf8ToBig5_test.go
func BenchmarkUtf8ToBig5(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		str := "你好，世界！"
		_,_ = toolbaktring.Utf8ToBig5(str)
	}
}
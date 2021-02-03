package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestUtf8ToGbk String/Global.go String/Utf8ToGbk.go String/Utf8ToGbk_test.go
func TestUtf8ToGbk(t *testing.T) {
	var toolbaktring String
	str := "你好，世界！"
	res1,_ := toolbaktring.Utf8ToGbk(str)
	fmt.Println(res1)
}

// go test -v -run TestUtf8ToGbk -bench=BenchmarkUtf8ToGbk -count=5 String/Global.go String/Utf8ToGbk.go String/Utf8ToGbk_test.go
func BenchmarkUtf8ToGbk(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		str := "你好，世界！"
		_,_ = toolbaktring.Utf8ToGbk(str)
	}
}
package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestBig5ToUtf8 String/Global.go String/Big5ToUtf8.go String/Big5ToUtf8_test.go String/Utf8ToBig5.go
func TestBig5ToUtf8(t *testing.T) {
	var toolbaktring String
	//str := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
	str := "你好，世界！"
	res1,_ := toolbaktring.Utf8ToBig5(str)
	res2,_ := toolbaktring.Big5ToUtf8(res1)
	fmt.Println(res1)
	fmt.Println(res2)
}

// go test -v -run TestBig5ToUtf8 -bench=BenchmarkBig5ToUtf8 -count=5 String/Global.go String/Big5ToUtf8.go String/Big5ToUtf8_test.go
func BenchmarkBig5ToUtf8(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		str := "你好，世界！"
		_,_ = toolbaktring.Big5ToUtf8(str)
	}
}
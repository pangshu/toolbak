package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGbkToUtf8 String/Global.go String/GbkToUtf8.go String/GbkToUtf8_test.go
func TestGbkToUtf8(t *testing.T) {
	var toolbaktring String
	//str := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
	str := string([]byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1})
	res1,_ := toolbaktring.GbkToUtf8(str)
	fmt.Println(res1)
}

// go test -v -run TestGbkToUtf8 -bench=BenchmarkGbkToUtf8 -count=5 String/Global.go String/GbkToUtf8.go String/GbkToUtf8_test.go
func BenchmarkGbkToUtf8(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		str := "你好，世界！"
		_,_ = toolbaktring.GbkToUtf8(str)
	}
}
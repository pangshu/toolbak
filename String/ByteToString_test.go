package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestByteToString String/Global.go String/ByteToString.go String/ByteToString_test.go
func TestByteToString(t *testing.T) {
	var toolbaktring String
	//str := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
	str := []byte("你好，世界！")
	res1 := toolbaktring.ByteToString(str)
	fmt.Println(res1)
}

// go test -v -run TestByteToString -bench=BenchmarkByteToString -count=5 String/Global.go String/ByteToString.go String/ByteToString_test.go
func BenchmarkByteToString(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		str := []byte("你好，世界！")
		_ = toolbaktring.ByteToString(str)
	}
}
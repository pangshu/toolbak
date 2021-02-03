package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestChunkSplit String/Global.go String/ChunkSplit.go String/ChunkSplit_test.go
func TestChunkSplit(t *testing.T) {
	var toolbaktring String
	//str := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
	str := "Yar?m kilo ?ay, yar?m kilo ?eker"
	res := toolbaktring.ChunkSplit(str, 4, "\r\n")
	fmt.Println(res)
	if len(res) == 0 {
		t.Error("ChunkSplit fail")
		return
	}
	fmt.Println(">>>>>>>>>>")
	res1 := toolbaktring.ChunkSplit(str, 5, "")
	fmt.Println(res1)
	res2 := toolbaktring.ChunkSplit("a", 4, "")
	fmt.Println(res2)
	res3 := toolbaktring.ChunkSplit("ab", 64, "")
	fmt.Println(res3)
	res4 := toolbaktring.ChunkSplit("abc", 1, "")
	fmt.Println(res4)
}

// go test -v -run TestChunkSplit -bench=BenchmarkChunkSplit -count=5 String/Global.go String/ChunkSplit.go String/ChunkSplit_test.go
func BenchmarkChunkSplit(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		str := "Yar?m kilo ?ay, yar?m kilo ?eker"
		_ = toolbaktring.ChunkSplit(str, 4, "\r\n")
	}
}
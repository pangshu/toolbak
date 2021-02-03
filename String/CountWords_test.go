package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestCountWords String/Global.go String/CountWords.go String/CountWords_test.go
func TestCountWords(t *testing.T) {
	var toolbaktring String
	str := "Hello 你好, World 世界！ 🍁🍃🍂🌰🍁🌿🌾🌼🌻hell中文"
	res1,res2 := toolbaktring.CountWords(str)
	fmt.Println(res1)
	fmt.Println(res2)
}

// go test -v -run TestCountWords -bench=BenchmarkCountWords -count=5 String/Global.go String/CountWords.go String/CountWords_test.go
func BenchmarkCountWords(t *testing.B) {
	t.ResetTimer()
	str := "Hello 你好, World 世界！"
	var toolbaktring String
	for i:=0; i< t.N; i++ {
		_,_ = toolbaktring.CountWords(str)
	}
}
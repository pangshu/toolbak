package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestClosestWord String/Global.go String/ClosestWord.go String/ClosestWord_test.go String/Levenshtein.go
func TestClosestWord(t *testing.T) {
	var toolbaktring String
	word := "hello,golang"
	search := []string{"hehe,php lang", "Hello,go language", "HeLlo,python!", "haha,java", "I`m going."}
	res1,res2 := toolbaktring.ClosestWord(word, search)
	fmt.Println(res1)
	fmt.Println(res2)
}

// go test -v -run TestClosestWord -bench=BenchmarkClosestWord -count=5 String/Global.go String/ClosestWord.go String/ClosestWord_test.go
func BenchmarkClosestWord(t *testing.B) {
	t.ResetTimer()
	word := "hello,golang"
	search := []string{"hehe,php lang", "Hello,go language", "HeLlo,python!", "haha,java", "I`m going."}
	var toolbaktring String
	for i:=0; i< t.N; i++ {
		_,_ = toolbaktring.ClosestWord(word, search)
	}
}
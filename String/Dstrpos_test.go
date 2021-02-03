package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDstrpos String/Global.go String/Dstrpos.go String/Dstrpos_test.go String/Strpos.go
func TestDstrpos(t *testing.T) {
	var toolbaktring String
	str := "Hello 你好, World 世界！"
	arr := []string{"he", "好", "world"}
	res1,res2 := toolbaktring.Dstrpos(str, arr, false)
	fmt.Println(res1)
	fmt.Println(res2)
}

// go test -v -run TestDstrpos -bench=BenchmarkDstrpos -count=5 String/Global.go String/Dstrpos.go String/Dstrpos_test.go
func BenchmarkDstrpos(t *testing.B) {
	t.ResetTimer()
	str := "Hello 你好, World 世界！"
	arr := []string{"he", "好", "world"}
	var toolbaktring String
	for i:=0; i< t.N; i++ {
		_,_ = toolbaktring.Dstrpos(str, arr, false)
	}
}
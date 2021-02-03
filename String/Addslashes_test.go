package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestAddslashes String/Global.go String/Addslashes.go String/Addslashes_test.go
func TestAddslashes(t *testing.T) {
	var toolbaktring String
	str := "Is your name O'reilly?"
	res1 := toolbaktring.Addslashes(str)
	fmt.Println(res1)
}

// go test -v -run TestAddslashes -bench=BenchmarkAddslashes -count=5 String/Global.go String/Addslashes.go String/Addslashes_test.go
func BenchmarkAddslashes(t *testing.B) {
	t.ResetTimer()
	var toolbaktring String
	str := "你好，世界！"
	for i:=0; i< t.N; i++ {
		_ = toolbaktring.Addslashes(str)
	}
}
package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDbcToSbc String/Global.go String/DbcToSbc.go String/DbcToSbc_test.go
func TestDbcToSbc(t *testing.T) {
	var toolbaktring String
	str := "Hello 你好, World 世界！ 🍁🍃🍂🌰🍁🌿🌾🌼🌻hell中文"
	res1 := toolbaktring.DbcToSbc(str)
	fmt.Println(res1)
}

// go test -v -run TestDbcToSbc -bench=BenchmarkDbcToSbc -count=5 String/Global.go String/DbcToSbc.go String/DbcToSbc_test.go
func BenchmarkDbcToSbc(t *testing.B) {
	t.ResetTimer()
	str := "Hello 你好, World 世界！"
	var toolbaktring String
	for i:=0; i< t.N; i++ {
		_ = toolbaktring.DbcToSbc(str)
	}
}
package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestContainsMap String/Global.go String/ContainsMap.go String/ContainsMap_test.go
func TestContainsMap(t *testing.T) {
	var toolbaktring String
	str := "Hello"  // 你好, World 世界！ 🍁🍃🍂🌰🍁🌿🌾🌼🌻hell中文
	arr := []string{"Hello", "好", "world"}
	res1 := toolbaktring.ContainsMap(str, arr)
	fmt.Println(res1)
}

// go test -v -run TestContainsMap -bench=BenchmarkContainsMap -count=5 String/Global.go String/ContainsMap.go String/ContainsMap_test.go
func BenchmarkContainsMap(t *testing.B) {
	t.ResetTimer()
	str := "Hello 你好, World 世界！"
	arr := []string{"he", "好", "world"}
	var toolbaktring String
	for i:=0; i< t.N; i++ {
		_ = toolbaktring.ContainsMap(str,arr)
	}
}
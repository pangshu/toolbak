package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDelEmoji String/Global.go String/DelEmoji.go String/DelEmoji_test.go
func TestDelEmoji(t *testing.T) {
	var toolbaktring String
	str := "Hello 你好, World 世界！ 🍁🍃🍂🌰🍁🌿🌾🌼🌻hell中文"
	res1 := toolbaktring.DelEmoji(str)
	fmt.Println(res1)
}

// go test -v -run TestDelEmoji -bench=BenchmarkDelEmoji -count=5 String/Global.go String/DelEmoji.go String/DelEmoji_test.go
func BenchmarkDelEmoji(t *testing.B) {
	t.ResetTimer()
	str := "Hello 你好, World 世界！"
	var toolbaktring String
	for i:=0; i< t.N; i++ {
		_ = toolbaktring.DelEmoji(str)
	}
}
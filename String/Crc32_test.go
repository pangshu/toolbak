package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestCrc32 String/Global.go String/Crc32.go String/Crc32_test.go
func TestCrc32(t *testing.T) {
	var toolbaktring String
	str := "Hello 你好, World 世界！ 🍁🍃🍂🌰🍁🌿🌾🌼🌻hell中文"
	res1 := toolbaktring.Crc32(str)
	fmt.Println(res1)
}

// go test -v -run TestCrc32 -bench=BenchmarkCrc32 -count=5 String/Global.go String/Crc32.go String/Crc32_test.go
func BenchmarkCrc32(t *testing.B) {
	t.ResetTimer()
	str := "Hello 你好, World 世界！"
	var toolbaktring String
	for i:=0; i< t.N; i++ {
		_ = toolbaktring.Crc32(str)
	}
}
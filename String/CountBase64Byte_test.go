package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestCountBase64Byte String/Global.go String/CountBase64Byte.go String/CountBase64Byte_test.go
func TestCountBase64Byte(t *testing.T) {
	var toolbaktring String
	str := "Hello 你好, World 世界！ 🍁🍃🍂🌰🍁🌿🌾🌼🌻hell中文"
	res1 := toolbaktring.CountBase64Byte(str)
	fmt.Println(res1)
}

// go test -v -run TestCountBase64Byte -bench=BenchmarkCountBase64Byte -count=5 String/Global.go String/CountBase64Byte.go String/CountBase64Byte_test.go
func BenchmarkCountBase64Byte(t *testing.B) {
	t.ResetTimer()
	str := "Hello 你好, World 世界！"
	var toolbaktring String
	for i:=0; i< t.N; i++ {
		_ = toolbaktring.CountBase64Byte(str)
	}
}
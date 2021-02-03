package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestChr String/Global.go String/Chr.go String/Chr_test.go
func TestChr(t *testing.T) {
	var toolbaktring String
	str := 65
	res1 := toolbaktring.Chr(str)
	fmt.Println(res1)
}

// go test -v -run TestChr -bench=BenchmarkChr -count=5 String/Global.go String/Chr.go String/Chr_test.go
func BenchmarkChr(t *testing.B) {
	t.ResetTimer()
	str := 65
	var toolbaktring String
	for i:=0; i< t.N; i++ {
		_ = toolbaktring.Chr(str)
	}
}
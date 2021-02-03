package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsBlank String/Global.go String/IsBlank.go String/IsBlank_test.go
func TestIsBlank(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.IsBlank("aa bb cc")
	fmt.Println(res1)
}

// go test -v -run TestIsBlank -bench=BenchmarkIsBlank -count=5 String/Global.go String/IsBlank.go String/IsBlank_test.go
func BenchmarkIsBlank(t *testing.B) {
	t.ResetTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.IsBlank("aa bb cc")
	}
}
package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsLower String/Global.go String/IsLower.go String/IsLower_test.go
func TestIsLower(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.IsLower("aa bb cc")
	fmt.Println(res1)
}

// go test -v -run TestIsLower -bench=BenchmarkIsLower -count=5 String/Global.go String/IsLower.go String/IsLower_test.go
func BenchmarkIsLower(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.IsLower("中文aa bb cc AaBbCcDd")
	}
}
package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsUpper String/Global.go String/IsUpper.go String/IsUpper_test.go
func TestIsUpper(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.IsUpper("中文aa bb cc AaBbCcDd")
	fmt.Println(res1)
}

// go test -v -run TestIsUpper -bench=BenchmarkIsUpper -count=5 String/Global.go String/IsUpper.go String/IsUpper_test.go
func BenchmarkIsUpper(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.IsUpper("中文aa bb cc AaBbCcDd")
	}
}
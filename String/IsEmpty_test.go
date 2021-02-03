package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsEmpty String/Global.go String/IsEmpty.go String/IsEmpty_test.go String/Trim.go
func TestIsEmpty(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.IsEmpty(" ")
	fmt.Println(res1)
}

// go test -v -run TestIsEmpty -bench=BenchmarkIsEmpty -count=5 String/Global.go String/IsEmpty.go String/IsEmpty_test.go
func BenchmarkIsEmpty(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.IsEmpty("aa bb cc")
	}
}
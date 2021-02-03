package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestFirstToUpper String/Global.go String/FirstToUpper.go String/FirstToUpper_test.go
func TestFirstToUpper(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.FirstToUpper("hello world. I love Shanghai!")
	fmt.Println(res1)
}

// go test -v -run TestFirstToUpper -bench=BenchmarkFirstToUpper -count=5 String/Global.go String/FirstToUpper.go String/FirstToUpper_test.go
func BenchmarkFirstToUpper(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.FirstToUpper("Hello world. I love Shanghai!")
	}
}
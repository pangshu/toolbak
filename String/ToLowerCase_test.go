package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestToLowerCase String/Global.go String/ToLowerCase.go String/ToLowerCase_test.go
func TestToLowerCase(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.ToLowerCase("aa bb cc AaBbCcDd", '_')
	fmt.Println(res1)
}

// go test -v -run TestToLowerCase -bench=BenchmarkToLowerCase -count=5 String/Global.go String/ToLowerCase.go String/ToLowerCase_test.go
func BenchmarkToLowerCase(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.ToLowerCase("aa bb cc AaBbCcDd", '_')
	}
}
package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestToMd5 String/Global.go String/ToMd5.go String/ToMd5_test.go
func TestToMd5(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.ToMd5("aa bb cc AaBbCcDd",16)
	res2 := toolbaktring.ToMd5("aa bb cc AaBbCcDd",32)
	fmt.Println(res1)
	fmt.Println(res2)
}

// go test -v -run TestToMd5 -bench=BenchmarkToMd5 -count=5 String/Global.go String/ToMd5.go String/ToMd5_test.go
func BenchmarkToMd5(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.ToMd5("aa bb cc AaBbCcDd")
	}
}
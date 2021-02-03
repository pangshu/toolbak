package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestSubstrCount String/Global.go String/SubstrCount.go String/SubstrCount_test.go
func TestSubstrCount(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.SubstrCount("aa bb cc AaBbCcDd", "a")
	fmt.Println(res1)
}

// go test -v -run TestSubstrCount -bench=BenchmarkSubstrCount -count=5 String/Global.go String/SubstrCount.go String/SubstrCount_test.go
func BenchmarkSubstrCount(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.SubstrCount("aa bb cc AaBbCcDd", "a")
	}
}
package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestSbcToDbc String/Global.go String/SbcToDbc.go String/SbcToDbc_test.go
func TestSbcToDbc(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.SbcToDbc("１２３４５６７８９ａｂｃ！")
	fmt.Println(res1)
}

// go test -v -run TestSbcToDbc -bench=BenchmarkSbcToDbc -count=5 String/Global.go String/SbcToDbc.go String/SbcToDbc_test.go
func BenchmarkSbcToDbc(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.SbcToDbc("１２３４５６７８９ａｂｃ！")
	}
}
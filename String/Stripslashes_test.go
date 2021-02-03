package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStripslashes String/Global.go String/Stripslashes.go String/Stripslashes_test.go
func TestStripslashes(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.Stripslashes("Who\\'s Bill Gates?")
	fmt.Println(res1)
}

// go test -v -run TestStripslashes -bench=BenchmarkStripslashes -count=5 String/Global.go String/Stripslashes.go String/Stripslashes_test.go
func BenchmarkStripslashes(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.Stripslashes("Who\\'s Bill Gates?")
	}
}
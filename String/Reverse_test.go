package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestReverse String/Global.go String/Reverse.go String/Reverse_test.go
func TestReverse(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.Reverse("Who's Bill Gates?")
	fmt.Println(res1)
}

// go test -v -run TestReverse -bench=BenchmarkReverse -count=5 String/Global.go String/Reverse.go String/Reverse_test.go
func BenchmarkReverse(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.Reverse("Who's Bill Gates?")
	}
}
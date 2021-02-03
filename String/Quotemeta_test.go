package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestQuotemeta String/Global.go String/Quotemeta.go String/Quotemeta_test.go
func TestQuotemeta(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.Quotemeta("Who's Bill Gates?")
	fmt.Println(res1)
}

// go test -v -run TestQuotemeta -bench=BenchmarkQuotemeta -count=5 String/Global.go String/Quotemeta.go String/Quotemeta_test.go
func BenchmarkQuotemeta(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.Quotemeta("Who's Bill Gates?")
	}
}
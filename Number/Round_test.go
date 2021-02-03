package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestRound Number/Global.go Number/Round.go Number/Round_test.go
func TestRound(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.Round(1.4432)
	fmt.Println(res1)
}

// go test -v -run TestRound -bench=BenchmarkRound -count=5 Number/Global.go Number/Round.go Number/Round_test.go
func BenchmarkRound(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Round(1)
	}
}
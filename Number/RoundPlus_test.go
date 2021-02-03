package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestRoundPlus Number/Global.go Number/RoundPlus.go Number/RoundPlus_test.go Number/Round.go
func TestRoundPlus(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.RoundPlus(1.4432, 3)
	fmt.Println(res1)
}

// go test -v -run TestRoundPlus -bench=BenchmarkRoundPlus -count=5 Number/Global.go Number/RoundPlus.go Number/RoundPlus_test.go Number/Round.go
func BenchmarkRoundPlus(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.RoundPlus(1, 3)
	}
}
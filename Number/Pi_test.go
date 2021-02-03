package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestPi Number/Global.go Number/Pi.go Number/Pi_test.go Number/SumFloat64.go
func TestPi(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.Pi()
	fmt.Println(res1)
}

// go test -v -run TestPi -bench=BenchmarkPi -count=5 Number/Global.go Number/Pi.go Number/Pi_test.go
func BenchmarkPi(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Pi()
	}
}
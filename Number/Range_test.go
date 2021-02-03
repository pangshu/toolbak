package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestRange Number/Global.go Number/Range.go Number/Range_test.go Number/AbsInt64.go
func TestRange(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.Range(5, 10)
	fmt.Println(res1)
}

// go test -v -run TestRange -bench=BenchmarkRange -count=5 Number/Global.go Number/Range.go Number/Range_test.go Number/AbsInt64.go
func BenchmarkRange(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Range(5, 10)
	}
}
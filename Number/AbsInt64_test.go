package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestAbsInt64 Number/Global.go Number/AbsInt64.go Number/AbsInt64_test.go
func TestAbsInt64(t *testing.T) {
	var toolNumber Number
	var num int64 = -123
	res1 := toolNumber.AbsInt64(num)
	fmt.Println(res1)
}

// go test -v -run TestAbsInt64 -bench=BenchmarkAbsInt64 -count=5 Number/Global.go Number/AbsInt64.go Number/AbsInt64_test.go
func BenchmarkAbsInt64(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	var num int64 = -123
	for i:=0; i< t.N; i++ {
		_ = toolNumber.AbsInt64(num)
	}
}
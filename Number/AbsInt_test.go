package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestAbsInt Number/Global.go Number/AbsInt.go Number/AbsInt_test.go
func TestAbsInt(t *testing.T) {
	var toolNumber Number
	num := -123
	res1 := toolNumber.AbsInt(num)
	fmt.Println(res1)
}

// go test -v -run TestAbsInt -bench=BenchmarkAbsInt -count=5 Number/Global.go Number/AbsInt.go Number/AbsInt_test.go
func BenchmarkAbsInt(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := -123
	for i:=0; i< t.N; i++ {
		_ = toolNumber.AbsInt(num)
	}
}
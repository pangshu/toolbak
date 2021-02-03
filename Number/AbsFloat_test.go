package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestAbsFloat Number/Global.go Number/AbsFloat.go Number/AbsFloat_test.go
func TestAbsFloat(t *testing.T) {
	var toolNumber Number
	num := -123.456
	res1 := toolNumber.AbsFloat(num)
	fmt.Println(res1)
}

// go test -v -run TestAbsFloat -bench=BenchmarkAbsFloat -count=5 Number/Global.go Number/AbsFloat.go Number/AbsFloat_test.go
func BenchmarkAbsFloat(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := -123.456
	for i:=0; i< t.N; i++ {
		_ = toolNumber.AbsFloat(num)
	}
}
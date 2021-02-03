package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestToFloat Number/Global.go Number/ToFloat.go Number/ToFloat_test.go
func TestToFloat(t *testing.T) {
	var toolNumber Number
	res1,_ := toolNumber.ToFloat(5.5)
	fmt.Println(res1)
}

// go test -v -run TestToFloat -bench=BenchmarkToFloat -count=5 Number/Global.go Number/ToFloat.go Number/ToFloat_test.go
func BenchmarkToFloat(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_,_ = toolNumber.ToFloat(5)
	}
}
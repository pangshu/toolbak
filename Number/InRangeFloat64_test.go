package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestInRangeFloat64 Number/Global.go Number/InRangeFloat64.go Number/InRangeFloat64_test.go
func TestInRangeFloat64(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.InRangeFloat64(2,1,3)
	fmt.Println(res1)
}

// go test -v -run TestInRangeFloat64 -bench=BenchmarkInRangeFloat64 -count=5 Number/Global.go Number/InRangeFloat64.go Number/InRangeFloat64_test.go
func BenchmarkInRangeFloat64(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.InRangeFloat64(2,1,3)
	}
}
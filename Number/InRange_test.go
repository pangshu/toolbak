package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestInRange Number/Global.go Number/InRange.go Number/InRange_test.go Number/InRangeInt.go Number/InRangeFloat32.go Number/InRangeFloat64.go
func TestInRange(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.InRange(2,1,3)
	fmt.Println(res1)
}

// go test -v -run TestInRange -bench=BenchmarkInRange -count=5 Number/Global.go Number/InRange.go Number/InRange_test.go Number/InRangeInt.go Number/InRangeFloat32.go Number/InRangeFloat64.go
func BenchmarkInRange(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.InRange(2,1,3)
	}
}
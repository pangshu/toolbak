package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestInRangeInt Number/Global.go Number/InRangeInt.go Number/InRangeInt_test.go
func TestInRangeInt(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.InRangeInt(2,1,3)
	fmt.Println(res1)
}

// go test -v -run TestInRangeInt -bench=BenchmarkInRangeInt -count=5 Number/Global.go Number/InRangeInt.go Number/InRangeInt_test.go
func BenchmarkInRangeInt(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.InRangeInt(2,1,3)
	}
}
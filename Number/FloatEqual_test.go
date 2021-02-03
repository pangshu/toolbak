package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestFloatEqual Number/Global.go Number/FloatEqual.go Number/FloatEqual_test.go
func TestFloatEqual(t *testing.T) {
	var toolNumber Number
	var f1, f2 float64
	f1 = 1.2345678
	f2 = 1.2345679
	res1 := toolNumber.FloatEqual(f1, f2, 6)
	fmt.Println(res1)
}

// go test -v -run TestFloatEqual -bench=BenchmarkFloatEqual -count=5 Number/Global.go Number/FloatEqual.go Number/FloatEqual_test.go
func BenchmarkFloatEqual(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	var f1, f2 float64
	f1 = 1.2345678
	f2 = 1.2345679
	for i:=0; i< t.N; i++ {
		_ = toolNumber.FloatEqual(f1, f2, 6)
	}
}
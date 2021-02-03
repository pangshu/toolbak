package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestInRangeFloat32 Number/Global.go Number/InRangeFloat32.go Number/InRangeFloat32_test.go
func TestInRangeFloat32(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.InRangeFloat32(2,1,3)
	fmt.Println(res1)
}

// go test -v -run TestInRangeFloat32 -bench=BenchmarkInRangeFloat32 -count=5 Number/Global.go Number/InRangeFloat32.go Number/InRangeFloat32_test.go
func BenchmarkInRangeFloat32(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.InRangeFloat32(2,1,3)
	}
}
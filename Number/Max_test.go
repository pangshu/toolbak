package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMax Number/Global.go Number/Max.go Number/Max_test.go Number/ToFloat.go
func TestMax(t *testing.T) {
	var toolNumber Number
	num := []interface{}{-1, 0, "18", true, nil, int8(1), int16(2), int32(3), int64(4), uint(5),
		uint8(6), uint16(7), uint32(8), uint64(9), float32(10.0), float64(11.1)}
	res1 := toolNumber.Max(num...)
	fmt.Println(res1)
}

// go test -v -run TestMax -bench=BenchmarkMax -count=5 Number/Global.go Number/Max.go Number/Max_test.go
func BenchmarkMax(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []interface{}{-1, 0, "18", true, nil, int8(1), int16(2), int32(3), int64(4), uint(5),
		uint8(6), uint16(7), uint32(8), uint64(9), float32(10.0), float64(11.1)}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Max(num...)
	}
}
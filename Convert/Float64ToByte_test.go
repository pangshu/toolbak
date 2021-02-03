package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestFloat64ToByte Convert/Global.go Convert/Float64ToByte.go Convert/Float64ToByte_test.go
func TestFloat64ToByte(t *testing.T) {
	var toolConvert Convert
	var str float64 = 12345.6
	res1 := toolConvert.Float64ToByte(str)
	fmt.Println(res1)
}

// go test -v -run TestFloat64ToByte -bench=BenchmarkFloat64ToByte -count=5 Convert/Global.go Convert/Float64ToByte.go Convert/Float64ToByte_test.go
func BenchmarkFloat64ToByte(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	var str float64 = 12345.6
	for i:=0; i< t.N; i++ {
		_ = toolConvert.Float64ToByte(str)
	}
}
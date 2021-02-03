package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestToFloat Convert/Global.go Convert/ToFloat.go Convert/ToFloat_test.go Convert/StrToFloat64.go Convert/StrToFloatStrict.go
func TestToFloat(t *testing.T) {
	var toolConvert Convert
	str := "1"
	res1 := toolConvert.ToFloat(str)
	fmt.Println(res1)
}

// go test -v -run TestToFloat -bench=BenchmarkToFloat -count=5 Convert/Global.go Convert/ToFloat.go Convert/ToFloat_test.go
func BenchmarkToFloat(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "1"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.ToFloat(str)
	}
}
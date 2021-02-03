package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestFloatToStr Convert/Global.go Convert/FloatToStr.go Convert/FloatToStr_test.go
func TestFloatToStr(t *testing.T) {
	var toolConvert Convert
	var str float64 = 12345.6789
	res1 := toolConvert.FloatToStr(str,1)
	fmt.Println(res1)
}

// go test -v -run TestFloatToStr -bench=BenchmarkFloatToStr -count=5 Convert/Global.go Convert/FloatToStr.go Convert/FloatToStr_test.go
func BenchmarkFloatToStr(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	var str float64 = 12345.6789
	for i:=0; i< t.N; i++ {
		_ = toolConvert.FloatToStr(str,1)
	}
}
package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIntToStr Convert/Global.go Convert/IntToStr.go Convert/IntToStr_test.go
func TestIntToStr(t *testing.T) {
	var toolConvert Convert
	var str int64 = 123456
	res1 := toolConvert.IntToStr(str)
	fmt.Println(res1)
}

// go test -v -run TestIntToStr -bench=BenchmarkIntToStr -count=5 Convert/Global.go Convert/IntToStr.go Convert/IntToStr_test.go
func BenchmarkIntToStr(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	var str int64 = 123456
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IntToStr(str)
	}
}
package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestToStr Convert/Global.go Convert/ToStr.go Convert/ToStr_test.go
func TestToStr(t *testing.T) {
	var toolConvert Convert
	str := 123.456
	res1 := toolConvert.ToStr(str)
	fmt.Println(res1)
}

// go test -v -run TestToStr -bench=BenchmarkToStr -count=5 Convert/Global.go Convert/ToStr.go Convert/ToStr_test.go
func BenchmarkToStr(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := 123.456
	for i:=0; i< t.N; i++ {
		_ = toolConvert.ToStr(str)
	}
}
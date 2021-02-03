package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestBoolToStr Convert/Global.go Convert/BoolToStr.go Convert/BoolToStr_test.go
func TestBoolToStr(t *testing.T) {
	var toolConvert Convert
	str := true
	res1 := toolConvert.BoolToStr(str)
	fmt.Println(res1)
}

// go test -v -run TestBoolToStr -bench=BenchmarkBoolToStr -count=5 Convert/Global.go Convert/BoolToStr.go Convert/BoolToStr_test.go
func BenchmarkBoolToStr(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := true
	for i:=0; i< t.N; i++ {
		_ = toolConvert.BoolToStr(str)
	}
}
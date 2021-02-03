package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToInt16 Convert/Global.go Convert/StrToInt16.go Convert/StrToInt16_test.go Convert/StrToIntStrict.go
func TestStrToInt16(t *testing.T) {
	var toolConvert Convert
	str := "123456"
	res1 := toolConvert.StrToInt16(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToInt16 -bench=BenchmarkStrToInt16 -count=5 Convert/Global.go Convert/StrToInt16.go Convert/StrToInt16_test.go Convert/StrToIntStrict.go
func BenchmarkStrToInt16(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToInt16(str)
	}
}
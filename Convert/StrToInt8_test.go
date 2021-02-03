package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToInt8 Convert/Global.go Convert/StrToInt8.go Convert/StrToInt8_test.go Convert/StrToIntStrict.go
func TestStrToInt8(t *testing.T) {
	var toolConvert Convert
	str := "123456"
	res1 := toolConvert.StrToInt8(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToInt8 -bench=BenchmarkStrToInt8 -count=5 Convert/Global.go Convert/StrToInt8.go Convert/StrToInt8_test.go Convert/StrToIntStrict.go
func BenchmarkStrToInt8(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToInt8(str)
	}
}
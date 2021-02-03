package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToInt32 Convert/Global.go Convert/StrToInt32.go Convert/StrToInt32_test.go Convert/StrToIntStrict.go
func TestStrToInt32(t *testing.T) {
	var toolConvert Convert
	str := "123456"
	res1 := toolConvert.StrToInt32(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToInt32 -bench=BenchmarkStrToInt32 -count=5 Convert/Global.go Convert/StrToInt32.go Convert/StrToInt32_test.go Convert/StrToIntStrict.go
func BenchmarkStrToInt32(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToInt32(str)
	}
}
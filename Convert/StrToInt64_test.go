package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToInt64 Convert/Global.go Convert/StrToInt64.go Convert/StrToInt64_test.go Convert/StrToIntStrict.go
func TestStrToInt64(t *testing.T) {
	var toolConvert Convert
	str := "123456"
	res1 := toolConvert.StrToInt64(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToInt64 -bench=BenchmarkStrToInt64 -count=5 Convert/Global.go Convert/StrToInt64.go Convert/StrToInt64_test.go Convert/StrToIntStrict.go
func BenchmarkStrToInt64(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToInt64(str)
	}
}
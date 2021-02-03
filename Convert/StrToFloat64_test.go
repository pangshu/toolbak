package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToFloat64 Convert/Global.go Convert/StrToFloat64.go Convert/StrToFloat64_test.go Convert/StrToFloatStrict.go
func TestStrToFloat64(t *testing.T) {
	var toolConvert Convert
	str := "123.456"
	res1 := toolConvert.StrToFloat64(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToFloat64 -bench=BenchmarkStrToFloat64 -count=5 Convert/Global.go Convert/StrToFloat64.go Convert/StrToFloat64_test.go Convert/StrToFloatStrict.go
func BenchmarkStrToFloat64(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123.456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToFloat64(str)
	}
}
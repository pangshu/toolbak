package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToUint64 Convert/Global.go Convert/StrToUint64.go Convert/StrToUint64_test.go Convert/StrToUintStrict.go
func TestStrToUint64(t *testing.T) {
	var toolConvert Convert
	str := "123456"
	res1 := toolConvert.StrToUint64(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToUint64 -bench=BenchmarkStrToUint64 -count=5 Convert/Global.go Convert/StrToUint64.go Convert/StrToUint64_test.go Convert/StrToUintStrict.go
func BenchmarkStrToUint64(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToUint64(str)
	}
}
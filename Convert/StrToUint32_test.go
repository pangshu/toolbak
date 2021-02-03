package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToUint32 Convert/Global.go Convert/StrToUint32.go Convert/StrToUint32_test.go Convert/StrToUintStrict.go
func TestStrToUint32(t *testing.T) {
	var toolConvert Convert
	str := "123456"
	res1 := toolConvert.StrToUint32(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToUint32 -bench=BenchmarkStrToUint32 -count=5 Convert/Global.go Convert/StrToUint32.go Convert/StrToUint32_test.go Convert/StrToUintStrict.go
func BenchmarkStrToUint32(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToUint32(str)
	}
}
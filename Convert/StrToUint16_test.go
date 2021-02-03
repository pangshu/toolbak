package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToUint16 Convert/Global.go Convert/StrToUint16.go Convert/StrToUint16_test.go Convert/StrToUintStrict.go
func TestStrToUint16(t *testing.T) {
	var toolConvert Convert
	str := "123456"
	res1 := toolConvert.StrToUint16(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToUint16 -bench=BenchmarkStrToUint16 -count=5 Convert/Global.go Convert/StrToUint16.go Convert/StrToUint16_test.go Convert/StrToUintStrict.go
func BenchmarkStrToUint16(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToUint16(str)
	}
}
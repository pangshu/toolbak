package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToUint8 Convert/Global.go Convert/StrToUint8.go Convert/StrToUint8_test.go Convert/StrToUintStrict.go
func TestStrToUint8(t *testing.T) {
	var toolConvert Convert
	str := "123456"
	res1 := toolConvert.StrToUint8(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToUint8 -bench=BenchmarkStrToUint8 -count=5 Convert/Global.go Convert/StrToUint8.go Convert/StrToUint8_test.go Convert/StrToUintStrict.go
func BenchmarkStrToUint8(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToUint8(str)
	}
}
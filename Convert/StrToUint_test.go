package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToUint Convert/Global.go Convert/StrToUint.go Convert/StrToUint_test.go Convert/StrToUintStrict.go
func TestStrToUint(t *testing.T) {
	var toolConvert Convert
	str := "123456"
	res1 := toolConvert.StrToUint(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToUint -bench=BenchmarkStrToUint -count=5 Convert/Global.go Convert/StrToUint.go Convert/StrToUint_test.go Convert/StrToUintStrict.go
func BenchmarkStrToUint(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToUint(str)
	}
}
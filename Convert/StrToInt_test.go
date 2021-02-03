package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToInt Convert/Global.go Convert/StrToInt.go Convert/StrToInt_test.go
func TestStrToInt(t *testing.T) {
	var toolConvert Convert
	str := "123456"
	res1 := toolConvert.StrToInt(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToInt -bench=BenchmarkStrToInt -count=5 Convert/Global.go Convert/StrToInt.go Convert/StrToInt_test.go
func BenchmarkStrToInt(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToInt(str)
	}
}
package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToUintStrict Convert/Global.go Convert/StrToUintStrict.go Convert/StrToUintStrict_test.go
func TestStrToUintStrict(t *testing.T) {
	var toolConvert Convert
	str := "123456"
	res1 := toolConvert.StrToUintStrict(str,32,true)
	fmt.Println(res1)
}

// go test -v -run TestStrToUintStrict -bench=BenchmarkStrToUintStrict -count=5 Convert/Global.go Convert/StrToUintStrict.go Convert/StrToUintStrict_test.go
func BenchmarkStrToUintStrict(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToUintStrict(str,32,true)
	}
}
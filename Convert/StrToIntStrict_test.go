package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToIntStrict Convert/Global.go Convert/StrToIntStrict.go Convert/StrToIntStrict_test.go
func TestStrToIntStrict(t *testing.T) {
	var toolConvert Convert
	str := "123456"
	res1 := toolConvert.StrToIntStrict(str,32,true)
	fmt.Println(res1)
}

// go test -v -run TestStrToIntStrict -bench=BenchmarkStrToIntStrict -count=5 Convert/Global.go Convert/StrToIntStrict.go Convert/StrToIntStrict_test.go
func BenchmarkStrToIntStrict(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToIntStrict(str,32,true)
	}
}
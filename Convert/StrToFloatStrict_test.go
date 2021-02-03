package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToFloatStrict Convert/Global.go Convert/StrToFloatStrict.go Convert/StrToFloatStrict_test.go
func TestStrToFloatStrict(t *testing.T) {
	var toolConvert Convert
	str := "123.456"
	res1 := toolConvert.StrToFloatStrict(str, 32, true)
	fmt.Println(res1)
}

// go test -v -run TestStrToFloatStrict -bench=BenchmarkStrToFloatStrict -count=5 Convert/Global.go Convert/StrToFloatStrict.go Convert/StrToFloatStrict_test.go
func BenchmarkStrToFloatStrict(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123.456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToFloatStrict(str, 32, true)
	}
}
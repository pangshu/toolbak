package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToBool Convert/Global.go Convert/StrToBool.go Convert/StrToBool_test.go
func TestStrToBool(t *testing.T) {
	var toolConvert Convert
	str := "true"
	res1 := toolConvert.StrToBool(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToBool -bench=BenchmarkStrToBool -count=5 Convert/Global.go Convert/StrToBool.go Convert/StrToBool_test.go
func BenchmarkStrToBool(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "true"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToBool(str)
	}
}
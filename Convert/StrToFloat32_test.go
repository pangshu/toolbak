package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToFloat32 Convert/Global.go Convert/StrToFloat32.go Convert/StrToFloat32_test.go Convert/StrToFloatStrict.go
func TestStrToFloat32(t *testing.T) {
	var toolConvert Convert
	str := "123.456"
	res1 := toolConvert.StrToFloat32(str)
	fmt.Println(res1)
}

// go test -v -run TestStrToFloat32 -bench=BenchmarkStrToFloat32 -count=5 Convert/Global.go Convert/StrToFloat32.go Convert/StrToFloat32_test.go Convert/StrToFloatStrict.go
func BenchmarkStrToFloat32(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123.456"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.StrToFloat32(str)
	}
}
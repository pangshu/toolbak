package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestBoolToInt Convert/Global.go Convert/BoolToInt.go Convert/BoolToInt_test.go
func TestBoolToInt(t *testing.T) {
	var toolConvert Convert
	str := true
	res1 := toolConvert.BoolToInt(str)
	fmt.Println(res1)
}

// go test -v -run TestBoolToInt -bench=BenchmarkBoolToInt -count=5 Convert/Global.go Convert/BoolToInt.go Convert/BoolToInt_test.go
func BenchmarkBoolToInt(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := true
	for i:=0; i< t.N; i++ {
		_ = toolConvert.BoolToInt(str)
	}
}
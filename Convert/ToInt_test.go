package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestToInt Convert/Global.go Convert/ToInt.go Convert/ToInt_test.go
func TestToInt(t *testing.T) {
	var toolConvert Convert
	str := "123"
	res1 := toolConvert.ToInt(str)
	fmt.Println(res1)
}

// go test -v -run TestToInt -bench=BenchmarkToInt -count=5 Convert/Global.go Convert/ToInt.go Convert/ToInt_test.go
func BenchmarkToInt(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "1"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.ToInt(str)
	}
}
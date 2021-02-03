package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsNumeric Convert/Global.go Convert/IsNumeric.go Convert/IsNumeric_test.go
func TestIsNumeric(t *testing.T) {
	var toolConvert Convert
	str := "123"
	res1 := toolConvert.IsNumeric(str)
	fmt.Println(res1)
}

// go test -v -run TestIsNumeric -bench=BenchmarkIsNumeric -count=5 Convert/Global.go Convert/IsNumeric.go Convert/IsNumeric_test.go
func BenchmarkIsNumeric(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := ""
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IsNumeric(str)
	}
}
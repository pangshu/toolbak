package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsInt Convert/Global.go Convert/IsInt.go Convert/IsInt_test.go
func TestIsInt(t *testing.T) {
	var toolConvert Convert
	str := "123456789"
	res1 := toolConvert.IsInt(str)
	fmt.Println(res1)
}

// go test -v -run TestIsInt -bench=BenchmarkIsInt -count=5 Convert/Global.go Convert/IsInt.go Convert/IsInt_test.go
func BenchmarkIsInt(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123456789"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IsInt(str)
	}
}
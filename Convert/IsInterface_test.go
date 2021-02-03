package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsInterface Convert/Global.go Convert/IsInterface.go Convert/IsInterface_test.go
func TestIsInterface(t *testing.T) {
	var toolConvert Convert
	str := "hi"
	res1 := toolConvert.IsInterface(str)
	fmt.Println(res1)
}

// go test -v -run TestIsInterface -bench=BenchmarkIsInterface -count=5 Convert/Global.go Convert/IsInterface.go Convert/IsInterface_test.go
func BenchmarkIsInterface(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "hi"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IsInterface(str)
	}
}
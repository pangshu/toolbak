package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsFloat Convert/Global.go Convert/IsFloat.go Convert/IsFloat_test.go
func TestIsFloat(t *testing.T) {
	var toolConvert Convert
	str := "123.12312"
	res1 := toolConvert.IsFloat(str)
	fmt.Println(res1)
}

// go test -v -run TestIsFloat -bench=BenchmarkIsFloat -count=5 Convert/Global.go Convert/IsFloat.go Convert/IsFloat_test.go
func BenchmarkIsFloat(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "123.123"
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IsFloat(str)
	}
}
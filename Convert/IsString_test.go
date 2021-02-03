package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsString Convert/Global.go Convert/IsString.go Convert/IsString_test.go Convert/Gettype.go
func TestIsString(t *testing.T) {
	var toolConvert Convert
	str := "123"
	res1 := toolConvert.IsString(str)
	fmt.Println(res1)
}

// go test -v -run TestIsString -bench=BenchmarkIsString -count=5 Convert/Global.go Convert/IsString.go Convert/IsString_test.go Convert/Gettype.go
func BenchmarkIsString(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := ""
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IsString(str)
	}
}
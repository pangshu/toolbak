package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsStruct Convert/Global.go Convert/IsStruct.go Convert/IsStruct_test.go Convert/Gettype.go
func TestIsStruct(t *testing.T) {
	var toolConvert Convert
	str := "123"
	res1 := toolConvert.IsStruct(str)
	fmt.Println(res1)
}

// go test -v -run TestIsStruct -bench=BenchmarkIsStruct -count=5 Convert/Global.go Convert/IsStruct.go Convert/IsStruct_test.go Convert/Gettype.go
func BenchmarkIsStruct(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := ""
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IsStruct(str)
	}
}
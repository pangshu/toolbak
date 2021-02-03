package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsEmpty Convert/Global.go Convert/IsEmpty.go Convert/IsEmpty_test.go
func TestIsEmpty(t *testing.T) {
	var toolConvert Convert
	str := ""
	res1 := toolConvert.IsEmpty(str)
	fmt.Println(res1)
}

// go test -v -run TestIsEmpty -bench=BenchmarkIsEmpty -count=5 Convert/Global.go Convert/IsEmpty.go Convert/IsEmpty_test.go
func BenchmarkIsEmpty(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := ""
	for i:=0; i< t.N; i++ {
		_ = toolConvert.IsEmpty(str)
	}
}
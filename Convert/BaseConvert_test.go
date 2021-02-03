package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestBaseConvert Convert/Global.go Convert/BaseConvert.go Convert/BaseConvert_test.go
func TestBaseConvert(t *testing.T) {
	var toolConvert Convert
	str := "1234567890"
	res1,err := toolConvert.BaseConvert(str, 10, 16)
	fmt.Println(res1)
	fmt.Println(err)
}

// go test -v -run TestBaseConvert -bench=BenchmarkBaseConvert -count=5 Convert/Global.go Convert/BaseConvert.go Convert/BaseConvert_test.go
func BenchmarkBaseConvert(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "1234567890"
	for i:=0; i< t.N; i++ {
		_,_ = toolConvert.BaseConvert(str, 10, 16)
	}
}
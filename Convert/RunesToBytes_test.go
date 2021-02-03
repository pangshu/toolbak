package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestRunesToBytes Convert/Global.go Convert/RunesToBytes.go Convert/RunesToBytes_test.go
func TestRunesToBytes(t *testing.T) {
	var toolConvert Convert
	str := []rune{'H', 'e', 'l', 'l', 'o', ' ', '世', '界'}
	res1 := toolConvert.RunesToBytes(str)
	fmt.Println(res1)
}

// go test -v -run TestRunesToBytes -bench=BenchmarkRunesToBytes -count=5 Convert/Global.go Convert/RunesToBytes.go Convert/RunesToBytes_test.go
func BenchmarkRunesToBytes(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := []rune{'H', 'e', 'l', 'l', 'o', ' ', '世', '界'}
	for i:=0; i< t.N; i++ {
		_ = toolConvert.RunesToBytes(str)
	}
}
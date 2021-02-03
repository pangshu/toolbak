package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestOctToDec Convert/Global.go Convert/OctToDec.go Convert/OctToDec_test.go
func TestOctToDec(t *testing.T) {
	var toolConvert Convert
	str := "726746425"
	res1,_ := toolConvert.OctToDec(str)
	fmt.Println(res1)
}

// go test -v -run TestOctToDec -bench=BenchmarkOctToDec -count=5 Convert/Global.go Convert/OctToDec.go Convert/OctToDec_test.go
func BenchmarkOctToDec(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := "726746425"
	for i:=0; i< t.N; i++ {
		_,_ = toolConvert.OctToDec(str)
	}
}
package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestInt64ToByte Convert/Global.go Convert/Int64ToByte.go Convert/Int64ToByte_test.go
func TestInt64ToByte(t *testing.T) {
	var toolConvert Convert
	var str int64 = 123456
	res1 := toolConvert.Int64ToByte(str)
	fmt.Println(res1)
}

// go test -v -run TestInt64ToByte -bench=BenchmarkInt64ToByte -count=5 Convert/Global.go Convert/Int64ToByte.go Convert/Int64ToByte_test.go
func BenchmarkInt64ToByte(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	var str int64 = 123456
	for i:=0; i< t.N; i++ {
		_ = toolConvert.Int64ToByte(str)
	}
}
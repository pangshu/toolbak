package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToTimestamp Date/Global.go Date/StrToTimestamp.go Date/StrToTimestamp_test.go Date/StrToTimestruct.go
func TestStrToTimestamp(t *testing.T) {
	var toolDate Date
	res1,_ := toolDate.StrToTimestamp("2019-07-11 10:11:23")
	fmt.Println(res1)
}

// go test -v -run TestStrToTimestamp -bench=BenchmarkStrToTimestamp -count=5 Date/Global.go Date/StrToTimestamp.go Date/StrToTimestamp_test.go
func BenchmarkStrToTimestamp(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_,_ = toolDate.StrToTimestamp("2019-07-11 10:11:23")
	}
}
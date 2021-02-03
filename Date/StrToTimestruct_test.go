package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStrToTimestruct Date/Global.go Date/StrToTimestruct.go Date/StrToTimestruct_test.go
func TestStrToTimestruct(t *testing.T) {
	var toolDate Date
	res1,_ := toolDate.StrToTimestruct("2019-07-11 10:11:23")
	fmt.Println(res1)
}

// go test -v -run TestStrToTimestruct -bench=BenchmarkStrToTimestruct -count=5 Date/Global.go Date/StrToTimestruct.go Date/StrToTimestruct_test.go
func BenchmarkStrToTimestruct(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_,_ = toolDate.StrToTimestruct("2019-07-11 10:11:23")
	}
}
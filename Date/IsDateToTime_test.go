package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsDateToTime Date/Global.go Date/IsDateToTime.go Date/IsDateToTime_test.go Date/StrToTimestamp.go Date/StrToTimestruct.go
func TestIsDateToTime(t *testing.T) {
	var toolDate Date
	res1,tt := toolDate.IsDateToTime("A2021-01-19")
	fmt.Println(res1)
	fmt.Println(tt)
}

// go test -v -run TestIsDateToTime -bench=BenchmarkIsDateToTime -count=5 Date/Global.go Date/IsDateToTime.go Date/IsDateToTime_test.go
func BenchmarkIsDateToTime(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_,_ = toolDate.IsDateToTime("2021-01-19")
	}
}
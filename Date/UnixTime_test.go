package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestUnixTime Date/Global.go Date/UnixTime.go Date/UnixTime_test.go
func TestUnixTime(t *testing.T) {
	var toolDate Date
	res1 := toolDate.UnixTime()
	fmt.Println(res1)
}

// go test -v -run TestUnixTime -bench=BenchmarkUnixTime -count=5 Date/Global.go Date/UnixTime.go Date/UnixTime_test.go
func BenchmarkUnixTime(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.UnixTime()
	}
}
package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestIsDate Date/Global.go Date/IsDate.go Date/IsDate_test.go
func TestIsDate(t *testing.T) {
	var toolDate Date
	res1 := toolDate.IsDate(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	fmt.Println(res1)
}

// go test -v -run TestIsDate -bench=BenchmarkIsDate -count=5 Date/Global.go Date/IsDate.go Date/IsDate_test.go
func BenchmarkIsDate(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.IsDate(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	}
}
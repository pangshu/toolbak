package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestUsleep Date/Global.go Date/Usleep.go Date/Usleep_test.go
func TestUsleep(t *testing.T) {
	var toolDate Date
	fmt.Println(">>>>>>")
	toolDate.Usleep(500000)
	fmt.Println(">>>>>>")
}

// go test -v -run TestUsleep -bench=BenchmarkUsleep -count=5 Date/Global.go Date/Usleep.go Date/Usleep_test.go
func BenchmarkUsleep(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		toolDate.Usleep(5)
	}
}
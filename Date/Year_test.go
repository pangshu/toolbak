package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestYear Date/Global.go Date/Year.go Date/Year_test.go
func TestYear(t *testing.T) {
	var toolDate Date
	res1 := toolDate.Year()
	fmt.Println(res1)
}

// go test -v -run TestYear -bench=BenchmarkYear -count=5 Date/Global.go Date/Year.go Date/Year_test.go
func BenchmarkYear(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.Year()
	}
}
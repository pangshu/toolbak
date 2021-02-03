package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestSecond Date/Global.go Date/Second.go Date/Second_test.go
func TestSecond(t *testing.T) {
	var toolDate Date
	res1 := toolDate.Second()
	fmt.Println(res1)
}

// go test -v -run TestSecond -bench=BenchmarkSecond -count=5 Date/Global.go Date/Second.go Date/Second_test.go
func BenchmarkSecond(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.Second()
	}
}
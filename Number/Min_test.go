package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMin Number/Global.go Number/Min.go Number/Min_test.go Number/ToFloat.go
func TestMin(t *testing.T) {
	var toolNumber Number
	num := []interface{}{-4, 0, 3, 9}
	res1 := toolNumber.Min(num...)
	fmt.Println(res1)
}

// go test -v -run TestMin -bench=BenchmarkMin -count=5 Number/Global.go Number/Min.go Number/Min_test.go Number/ToFloat.go
func BenchmarkMin(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []interface{}{-4, 0, 3, 9}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Min(num...)
	}
}
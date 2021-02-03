package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMaxInt Number/Global.go Number/MaxInt.go Number/MaxInt_test.go
func TestMaxInt(t *testing.T) {
	var toolNumber Number
	num := []int{-4, 0, 3, 9}
	res1 := toolNumber.MaxInt(num...)
	fmt.Println(res1)
}

// go test -v -run TestMaxInt -bench=BenchmarkMaxInt -count=5 Number/Global.go Number/MaxInt.go Number/MaxInt_test.go
func BenchmarkMaxInt(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []int{-4, 0, 3, 9}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.MaxInt(num...)
	}
}
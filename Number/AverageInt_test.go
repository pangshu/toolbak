package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestAverageInt Number/Global.go Number/AverageInt.go Number/AverageInt_test.go Number/SumInt.go
func TestAverageInt(t *testing.T) {
	var toolNumber Number
	num := []int{1,2,3,4,5,6,7,8,9,10}
	res1 := toolNumber.AverageInt(num)
	fmt.Println(res1)
}

// go test -v -run TestAverageInt -bench=BenchmarkAverageInt -count=5 Number/Global.go Number/AverageInt.go Number/AverageInt_test.go Number/SumInt.go
func BenchmarkAverageInt(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []int{1,2,3,4,5,6,7,8,9,10}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.AverageInt(num)
	}
}
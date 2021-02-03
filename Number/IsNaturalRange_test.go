package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsNaturalRange Number/Global.go Number/IsNaturalRange.go Number/IsNaturalRange_test.go Number/Range.go Number/AbsInt64.go
func TestIsNaturalRange(t *testing.T) {
	var toolNumber Number
	arr1 := []int{0, 1, 2, 3}
	//arr2 := []int{0, 3, 1, 2}
	res1 := toolNumber.IsNaturalRange(arr1, false)
	fmt.Println(res1)
}

// go test -v -run TestIsNaturalRange -bench=BenchmarkIsNaturalRange -count=5 Number/Global.go Number/IsNaturalRange.go Number/IsNaturalRange_test.go Number/Range.go Number/AbsInt64.go
func BenchmarkIsNaturalRange(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	arr1 := []int{1, 2, 3}
	//arr2 := []int{0, 3, 1, 2}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.IsNaturalRange(arr1, false)
	}
}
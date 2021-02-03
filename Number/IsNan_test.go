package Number

import (
	"fmt"
	"math"
	"testing"
)

// 测试命令: go test -v -run TestIsNan Number/Global.go Number/IsNan.go Number/IsNan_test.go
func TestIsNan(t *testing.T) {
	var toolNumber Number
	arr1 := math.Acos(1.01)
	res1 := toolNumber.IsNan(arr1)
	fmt.Println(res1)
}

// go test -v -run TestIsNan -bench=BenchmarkIsNan -count=5 Number/Global.go Number/IsNan.go Number/IsNan_test.go
func BenchmarkIsNan(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	arr1 := math.Acos(1.01)
	for i:=0; i< t.N; i++ {
		_ = toolNumber.IsNan(arr1)
	}
}
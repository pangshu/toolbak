package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsWhole Number/Global.go Number/IsWhole.go Number/IsWhole_test.go
func TestIsWhole(t *testing.T) {
	var toolNumber Number
	num := 5.0
	res1 := toolNumber.IsWhole(num)
	fmt.Println(res1)
}

// go test -v -run TestIsWhole -bench=BenchmarkIsWhole -count=5 Number/Global.go Number/IsWhole.go Number/IsWhole_test.go
func BenchmarkIsWhole(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := 5.5
	for i:=0; i< t.N; i++ {
		_ = toolNumber.IsWhole(num)
	}
}
package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestAverage Number/Global.go Number/Average.go Number/Average_test.go
func TestAverage(t *testing.T) {
	var toolNumber Number
	num := []interface{}{1,2.2,3.3,4,5,6,7,8,9,10}
	res1 := toolNumber.Average(num)
	fmt.Println(res1)
}

// go test -v -run TestAverage -bench=BenchmarkAverage -count=5 Number/Global.go Number/Average.go Number/Average_test.go
func BenchmarkAverage(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []interface{}{1,2,3,4,5,6,7,8,9,10}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Average(num)
	}
}
package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestOrd String/Global.go String/Ord.go String/Ord_test.go
func TestOrd(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.Ord("a")
	fmt.Println(res1)
}

// go test -v -run TestOrd -bench=BenchmarkOrd -count=5 String/Global.go String/Ord.go String/Ord_test.go
func BenchmarkOrd(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.Ord("a")
	}
}
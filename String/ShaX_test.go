package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestShaX String/Global.go String/ShaX.go String/ShaX_test.go
func TestShaX(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.ShaX("apple", 2)
	fmt.Println(res1)
}

// go test -v -run TestShaX -bench=BenchmarkShaX -count=5 String/Global.go String/ShaX.go String/ShaX_test.go
func BenchmarkShaX(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.ShaX("aa bb cc AaBbCcDd", 256)
	}
}
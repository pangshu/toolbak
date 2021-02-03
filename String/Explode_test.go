package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestExplode String/Global.go String/Explode.go String/Explode_test.go
func TestExplode(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.Explode("Hello world. I love Shanghai!", "")
	fmt.Println(res1)
}

// go test -v -run TestExplode -bench=BenchmarkExplode -count=5 String/Global.go String/Explode.go String/Explode_test.go
func BenchmarkExplode(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.Explode("Hello world. I love Shanghai!"," ")
	}
}
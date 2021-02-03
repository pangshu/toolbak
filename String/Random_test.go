package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestRandom String/Global.go String/Random.go String/Random_test.go
func TestRandom(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.Random(8,0)
	fmt.Println(res1)
	res1 = toolbaktring.Random(8,1)
	fmt.Println(res1)
	res1 = toolbaktring.Random(8,2)
	fmt.Println(res1)
	res1 = toolbaktring.Random(8,3)
	fmt.Println(res1)
	res1 = toolbaktring.Random(8,4)
	fmt.Println(res1)

}

// go test -v -run TestRandom -bench=BenchmarkRandom -count=5 String/Global.go String/Random.go String/Random_test.go
func BenchmarkRandom(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.Random(8,1)
	}
}
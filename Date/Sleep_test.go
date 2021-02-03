package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestSleep Date/Global.go Date/Sleep.go Date/Sleep_test.go
func TestSleep(t *testing.T) {
	var toolDate Date
	fmt.Println(time.Now().Unix())
	toolDate.Sleep(5)
	fmt.Println(time.Now().Unix())
	//fmt.Println(res1)
}

// go test -v -run TestSleep -bench=BenchmarkSleep -count=5 Date/Global.go Date/Sleep.go Date/Sleep_test.go
func BenchmarkSleep(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		toolDate.Sleep(1)
	}
}
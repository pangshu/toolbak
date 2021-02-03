package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestVersionCompare String/Global.go String/VersionCompare.go String/VersionCompare_test.go
func TestVersionCompare(t *testing.T) {
	var toolbaktring String
	v1 := "v1.1.1"
	v2 := "v1.2.10b"
	res1 := toolbaktring.VersionCompare(v1, v2, "<")
	fmt.Println(res1)
}

// go test -v -run TestVersionCompare -bench=BenchmarkVersionCompare -count=5 String/Global.go String/VersionCompare.go String/VersionCompare_test.go
func BenchmarkVersionCompare(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		v1 := "v1.1.1"
		v2 := "v1.2.10b"
		_ = toolbaktring.VersionCompare(v1, v2, "<")
	}
}
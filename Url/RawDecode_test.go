package Url

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestRawDecode Url/Global.go Url/RawDecode.go Url/RawDecode_test.go
func TestRawDecode(t *testing.T) {
	var toolUrl Url
	str := `foo%20bar%40baz`
	res1,_ := toolUrl.RawDecode(str)
	fmt.Println(res1)
}

// go test -v -run TestRawDecode -bench=BenchmarkRawDecode -count=5 Url/Global.go Url/RawDecode.go Url/RawDecode_test.go  Url/IsUrl.go
func BenchmarkRawDecode(t *testing.B) {
	t.ResetTimer()
	str := `foo%20bar%40baz`
	var toolUrl Url
	for i:=0; i< t.N; i++ {
		_,_ = toolUrl.RawDecode(str)
	}
}
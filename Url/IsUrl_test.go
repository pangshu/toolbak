package Url

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsUrl Url/Global.go Url/IsUrl.go Url/IsUrl_test.go
func TestIsUrl(t *testing.T) {
	var toolUrl Url
	str := "http://login.localhost:3000"
	res1 := toolUrl.IsUrl(str)
	fmt.Println(res1)
}

// go test -v -run TestIsUrl -bench=BenchmarkIsUrl -count=5 Url/Global.go Url/IsUrl.go Url/IsUrl_test.go
func BenchmarkIsUrl(t *testing.B) {
	t.ResetTimer()
	str := "http://login.localhost:3000"
	var toolUrl Url
	for i:=0; i< t.N; i++ {
		_ = toolUrl.IsUrl(str)
	}
}
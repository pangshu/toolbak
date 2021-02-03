package Url

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsUrlExists Url/Global.go Url/IsUrlExists.go Url/IsUrlExists_test.go  Url/IsUrl.go
func TestIsUrlExists(t *testing.T) {
	var toolUrl Url
	str := "http://baidu.com"
	res1 := toolUrl.IsUrlExists(str)
	fmt.Println(res1)
}

// go test -v -run TestIsUrlExists -bench=BenchmarkIsUrlExists -count=5 Url/Global.go Url/IsUrlExists.go Url/IsUrlExists_test.go  Url/IsUrl.go
func BenchmarkIsUrlExists(t *testing.B) {
	t.ResetTimer()
	str := "http://login.localhost:3000"
	var toolUrl Url
	for i:=0; i< t.N; i++ {
		_ = toolUrl.IsUrlExists(str)
	}
}
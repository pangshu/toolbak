package Url

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDomain Url/Global.go Url/Domain.go Url/Domain_test.go Url/FormatUrl.go
func TestDomain(t *testing.T) {
	var toolUrl Url
	str := "http://login.localhost:3000"
	res1 := toolUrl.Domain(str)
	fmt.Println(res1)
}

// go test -v -run TestDomain -bench=BenchmarkDomain -count=5 Url/Global.go Url/Domain.go Url/Domain_test.go
func BenchmarkDomain(t *testing.B) {
	t.ResetTimer()
	str := "http://login.localhost:3000"
	var toolUrl Url
	for i:=0; i< t.N; i++ {
		_ = toolUrl.Domain(str)
	}
}
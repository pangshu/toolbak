package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestPwd Os/Global.go Os/Pwd.go Os/Pwd_test.go
func TestPwd(t *testing.T) {
	var toolOs Os
	res1 := toolOs.Pwd()
	fmt.Println(res1)
}

// go test -v -run TestPwd -bench=BenchmarkPwd -count=5 Os/Global.go Os/Pwd.go Os/Pwd_test.go
func BenchmarkPwd(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.Pwd()
	}
}
package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetPwd Os/Global.go Os/GetPwd.go Os/GetPwd_test.go
func TestGetPwd(t *testing.T) {
	var toolOs Os
	res1,_ := toolOs.GetPwd()
	fmt.Println(res1)
}

// go test -v -run TestGetPwd -bench=BenchmarkGetPwd -count=5 Os/Global.go Os/GetPwd.go Os/GetPwd_test.go
func BenchmarkGetPwd(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_,_ = toolOs.GetPwd()
	}
}
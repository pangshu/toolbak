package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestExec Os/Global.go Os/Exec.go Os/Exec_test.go
func TestExec(t *testing.T) {
	var toolOs Os
	cmd := " ls -a -h"
	res1,res2,res3 := toolOs.Exec(cmd)
	fmt.Println(res1)
	fmt.Println(string(res2))
	fmt.Println(string(res3))
}

// go test -v -run TestExec -bench=BenchmarkExec -count=5 Os/Global.go Os/Exec.go Os/Exec_test.go
func BenchmarkExec(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	cmd := " ls -a -h"
	for i:=0; i< t.N; i++ {
		_,_,_ = toolOs.Exec(cmd)
	}
}
package Os

import (
	"fmt"
	"os/user"
	"testing"
	"github.com/pangshu/toolbak/Convert"
)

// 测试命令: go test -v -run TestChown Os/Global.go Os/Chown.go Os/Chown_test.go
func TestChown(t *testing.T) {
	var toolOs Os
	var toolConvert Convert.Convert
	usr, _ := user.Current()
	uid := toolConvert.ToInt(usr.Uid)
	guid := toolConvert.ToInt(usr.Gid)
	fmt.Println(uid, " --- ", guid)
	file := "/www/golang/src/toolbak/test.go"
	res1 := toolOs.Chown(file, uid, guid)
	fmt.Println(res1)
}

// go test -v -run TestChown -bench=BenchmarkChown -count=5 Os/Global.go Os/Chown.go Os/Chown_test.go
func BenchmarkChown(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	var toolConvert Convert.Convert
	usr, _ := user.Current()
	uid := toolConvert.ToInt(usr.Uid)
	guid := toolConvert.ToInt(usr.Gid)
	file := "/www/golang/src/toolbak/test.go"
	for i:=0; i< t.N; i++ {
		_ = toolOs.Chown(file, uid, guid)
	}
}
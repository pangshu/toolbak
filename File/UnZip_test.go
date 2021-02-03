package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestUnZip File/Global.go File/UnZip.go File/UnZip_test.go File/AbsPath.go File/Dirname.go File/IsExist.go File/Mkdir.go File/FileTree.go File/IsFile.go File/IsDir.go
func TestUnZip(t *testing.T) {
	var toolFile File
	res1,_ := toolFile.UnZip("/www/golang/src/v5/main222.zip", "/www/golang/src/v5")
	fmt.Println(res1)
}

// go test -v -run TestUnZip -bench=BenchmarkUnZip -count=5 File/Global.go File/UnZip.go File/UnZip_test.go
func BenchmarkUnZip(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.UnZip("/www/golang/src/v5/main222.go","/www/golang/src/v5/")
	}
}
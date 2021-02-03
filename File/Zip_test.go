package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestZip File/Global.go File/Zip.go File/Zip_test.go File/AbsPath.go File/Dirname.go File/IsExist.go File/Mkdir.go File/FileTree.go File/IsFile.go File/IsDir.go
func TestZip(t *testing.T) {
	var toolFile File
	res1,_ := toolFile.Zip("/www/golang/src/v5/main111.zip", "/www/golang/src/v5/main111.go")
	fmt.Println(res1)
}

// go test -v -run TestZip -bench=BenchmarkZip -count=5 File/Global.go File/Zip.go File/Zip_test.go
func BenchmarkZip(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.Zip("/www/golang/src/v5/main222.go","/www/golang/src/v5/")
	}
}
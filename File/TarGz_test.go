package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestTarGz File/Global.go File/TarGz.go File/TarGz_test.go File/AbsPath.go File/Dirname.go File/IsExist.go File/Mkdir.go File/FileTree.go File/IsFile.go File/IsDir.go
func TestTarGz(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main222.go"
	res1,_ := toolFile.TarGz(form, "/www/golang/src/v5/main222.tar.gz")
	fmt.Println(res1)
}

// go test -v -run TestTarGz -bench=BenchmarkTarGz -count=5 File/Global.go File/TarGz.go File/TarGz_test.go
func BenchmarkTarGz(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.TarGz(form,"/www/golang/src/v5/main222.tar.gz")
	}
}
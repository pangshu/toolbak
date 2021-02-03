package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestUnTarGz File/Global.go File/UnTarGz.go File/UnTarGz_test.go File/AbsPath.go File/Dirname.go File/IsExist.go File/Mkdir.go File/FileTree.go File/IsFile.go File/IsDir.go
func TestUnTarGz(t *testing.T) {
	var toolFile File
	res1,err := toolFile.UnTarGz("/www/golang/src/v5/main222.tar.gz", "/www/golang/src/v5/")
	fmt.Println(">>>>>>>>>>>>>>")
	fmt.Println(res1)
	fmt.Println(err)
	fmt.Println(">>>>>>>>>>>>>>")
}

// go test -v -run TestUnTarGz -bench=BenchmarkUnTarGz -count=5 File/Global.go File/UnTarGz.go File/UnTarGz_test.go
func BenchmarkUnTarGz(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.UnTarGz("/www/golang/src/v5/main222.tar.gz","/www/golang/src/v5/")
	}
}
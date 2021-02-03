package File

import (
	"path/filepath"
	"regexp"
	"strings"
)

// FormatPath 格式化路径
func (toolFile *File)FormatPath(filePath string) string {
	if filePath == "" {
		return ""
	}

	//替换特殊字符
	filePath = strings.NewReplacer(`|`, "", `:`, "", `<`, "", `>`, "", `?`, "", `\`, "/").Replace(filePath)

	//替换连续斜杠, 连续的"//"或"\\"或"\/"或"/\"
	RegFormatDir := regexp.MustCompile(`[\/]{2,}`)

	filePath = RegFormatDir.ReplaceAllString(filePath, "/")

	dir := toolFile.FormatDir(filepath.Dir(filePath))

	if dir == `./` && strings.Index(filePath, dir) != -1 {
		return filePath
	}

	return dir + filepath.Base(filePath)
}
//
//func TestFormatPath(t *testing.T) {
//	p1 := `/usr\bin\\golang//fmt/\test\/hehe`
//	p2 := `/usr|///tmp:\\\123/\abc<|\hello>\/%world?\\how$\\are\@#test.png`
//	p3 := `test.log`
//	p4 := `./test.log`
//
//	res1 := FormatPath(p1)
//	res2 := FormatPath(p2)
//	res3 := FormatPath(p3)
//	res4 := FormatPath(p4)
//
//	if res1 != `/usr/bin/golang/fmt/test/hehe` {
//		t.Error("FormatPath fail")
//		return
//	}
//	if res2 != `/usr/tmp/123/abc/hello/%world/how$/are/@#test.png` {
//		t.Error("FormatPath fail")
//		return
//	}
//	if res3 != `test.log` {
//		t.Error("FormatPath fail")
//		return
//	}
//	if res4 != `./test.log` {
//		t.Error("FormatPath fail")
//		return
//	}
//
//	FormatPath("")
//}


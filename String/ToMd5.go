package String

import (
	"crypto/md5"
	"fmt"
)

// 字符串生成md5
func (*String)ToMd5(str string, length ...int) string {
	if len(str) == 0 {
		return ""
	}
	var num int
	if len(length) == 0 {
		num = 32
	} else {
		if length[0] != 16 && length[0] != 32 {
			num = 32
		} else {
			num = length[0]
		}
	}
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)

	if num == 16 {
		return md5str[8:24]
	} else {
		return md5str
	}
	//return md5str
}
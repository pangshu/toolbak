package Convert

import "fmt"

// Gettype 获取变量类型.
func (*Convert)Gettype(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

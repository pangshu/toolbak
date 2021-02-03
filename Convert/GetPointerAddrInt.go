package Convert

import "fmt"

// GetPointerAddrInt 获取变量指针地址整型值.variable为变量.
func (toolConvert *Convert)GetPointerAddrInt(variable interface{}) int64 {
	res, _ := toolConvert.HexToDec(fmt.Sprintf("%p", &variable))
	return res
}

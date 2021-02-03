package Convert

import "fmt"

// IntToStr 将整数转换为字符串.
func (*Convert)IntToStr(val interface{}) string {
	switch val.(type) {
	// Integers
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	// Type is not integers, return empty string
	default:
		return ""
	}
}

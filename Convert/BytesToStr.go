package Convert

import "unsafe"

// BytesToStr 将字节切片转换为字符串.
// 零拷贝,不安全.效率是string([]byte{})的百倍以上,且转换量越大效率优势越明显.
func (*Convert)BytesToStr(val []byte) string {
	return *(*string)(unsafe.Pointer(&val))
}

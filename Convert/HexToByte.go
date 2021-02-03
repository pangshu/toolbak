package Convert

import "encoding/hex"

// HexToByte 16进制字符串转字节切片.
func (*Convert)HexToByte(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}

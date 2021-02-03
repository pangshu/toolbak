package Convert

import "encoding/hex"

// ByteToHex 字节切片转16进制字符串.
func (*Convert)ByteToHex(val []byte) string {
	return hex.EncodeToString(val)
}

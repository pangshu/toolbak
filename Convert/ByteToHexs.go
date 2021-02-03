package Convert

import "encoding/hex"

// ByteToHexs 字节切片转16进制切片.
func (*Convert)ByteToHexs(val []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(val)))
	hex.Encode(dst, val)
	return dst
}

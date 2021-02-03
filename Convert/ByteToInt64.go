package Convert

import "encoding/binary"

// ByteToInt64 字节切片转64位整型.
func (*Convert)ByteToInt64(val []byte) int64 {
	return int64(binary.BigEndian.Uint64(val))
}

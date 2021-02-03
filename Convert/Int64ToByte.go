package Convert

import "encoding/binary"

// Int64ToByte 64位整型转字节切片.
func (*Convert)Int64ToByte(val int64) []byte {
	res := make([]byte, 8)
	binary.BigEndian.PutUint64(res, uint64(val))

	return res
}

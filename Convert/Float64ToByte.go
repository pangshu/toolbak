package Convert

import (
	"encoding/binary"
	"math"
)

// Float64ToByte 64位浮点数转字节切片.
func (*Convert)Float64ToByte(val float64) []byte {
	bits := math.Float64bits(val)
	res := make([]byte, 8)
	binary.LittleEndian.PutUint64(res, bits)

	return res
}

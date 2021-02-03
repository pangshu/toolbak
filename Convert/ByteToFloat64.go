package Convert

import (
	"encoding/binary"
	"math"
)

// ByteToFloat64 字节切片转64位浮点数.
func (*Convert)ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

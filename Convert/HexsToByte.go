package Convert

import "encoding/hex"

// HexsToByte 16进制切片转byte切片.
func (*Convert)HexsToByte(val []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(val)))
	_, err := hex.Decode(dst, val)

	if err != nil {
		return nil
	}
	return dst
}

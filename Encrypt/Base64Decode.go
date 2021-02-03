package Encrypt

import "encoding/base64"

// Base64Decode 对使用 MIME base64 编码的数据进行解码.
func Base64Decode(str []byte) ([]byte, error) {
	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(str)))
	n, err := base64.StdEncoding.Decode(dbuf, str)
	return dbuf[:n], err
}

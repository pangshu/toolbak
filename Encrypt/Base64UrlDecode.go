package Encrypt

import (
	"bytes"
	"encoding/base64"
)

// Base64UrlDecode url安全的Base64Decode.
func Base64UrlDecode(data []byte) ([]byte, error) {
	var missing = (4 - len(data)%4) % 4
	data = append(data, bytes.Repeat([]byte("="), missing)...)

	dbuf := make([]byte, base64.URLEncoding.DecodedLen(len(data)))
	n, err := base64.URLEncoding.Decode(dbuf, data)
	return dbuf[:n], err
}

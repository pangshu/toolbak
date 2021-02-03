package Encrypt

import (
	"bytes"
	"encoding/base64"
)

// Base64UrlSafeEncode url安全的Base64Encode,没有'/'和'+'及结尾的'=' .
func Base64UrlEncode(source []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(source)))
	base64.StdEncoding.Encode(buf, source)

	// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
	buf = bytes.Replace(buf, []byte("/"), []byte("_"), -1)
	buf = bytes.Replace(buf, []byte("+"), []byte("-"), -1)
	buf = bytes.Replace(buf, []byte("="), []byte(""), -1)

	return buf
}

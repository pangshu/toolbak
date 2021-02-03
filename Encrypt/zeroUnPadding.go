package Encrypt

import "bytes"

// zeroUnPadding PKCS7-0拆解.
func zeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

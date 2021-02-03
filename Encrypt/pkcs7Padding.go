package Encrypt

import "bytes"

// pkcs7Padding PKCS7填充.
// cipherText为密文;blockSize为分组长度;isZero是否零填充.
func pkcs7Padding(cipherText []byte, blockSize int, isZero bool) []byte {
	textLength := len(cipherText)
	if cipherText == nil || textLength == 0 || blockSize <= 0 {
		return nil
	}

	var textByte []byte
	padding := blockSize - textLength%blockSize
	if isZero {
		textByte = bytes.Repeat([]byte{0}, padding)
	} else {
		textByte = bytes.Repeat([]byte{byte(padding)}, padding)
	}

	return append(cipherText, textByte...)
}

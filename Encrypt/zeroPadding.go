package Encrypt

// zeroPadding PKCS7使用0填充.
func zeroPadding(cipherText []byte, blockSize int) []byte {
	return pkcs7Padding(cipherText, blockSize, true)
}

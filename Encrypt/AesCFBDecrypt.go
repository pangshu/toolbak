package Encrypt

// AesCFBDecrypt AES-CFB密文反馈(Cipher feedback)模式解密.
// cipherText为密文;key为密钥,长16/24/32.
func AesCFBDecrypt(cipherText, key []byte) ([]byte, error) {
	return AesDecrypt(cipherText, key, "CFB", -1)
}

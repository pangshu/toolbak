package Encrypt

// AesOFBDecrypt AES-OFB输出反馈(Output feedback)模式解密.
// cipherText为密文;key为密钥,长16/24/32.
func AesOFBDecrypt(cipherText, key []byte) ([]byte, error) {
	return AesDecrypt(cipherText, key, "OFB", -1)
}

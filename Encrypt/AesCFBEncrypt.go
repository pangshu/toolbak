package Encrypt

// AesCFBEncrypt AES-CFB密文反馈(Cipher feedback)模式加密.适合对流数据加密.
// clearText为明文;key为密钥,长16/24/32.
func AesCFBEncrypt(clearText, key []byte) ([]byte, error) {
	return AesEncrypt(clearText, key, "CFB", -1)
}
